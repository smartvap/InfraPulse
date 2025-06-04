package server

import (
	"context"
	"crypto/tls"
	"github.com/clarechu/infrapulse/pkg/server/router"
	"k8s.io/klog/v2"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Bootstrap interface {
	ListenAndServe()
	Stop()
}

type CMDB struct {
	ctx       context.Context
	cancel    context.CancelFunc
	server    router.Server
	tls       *TLSOptions
	port      int32
	proxyPort int32
}

// TLSOptions holds the TLS options.
type TLSOptions struct {
	Config   *tls.Config
	CertFile string
	KeyFile  string
}

func (c *CMDB) Stop() {
	//TODO implement me
	panic("implement me")
}

// ListenAndServe runs the kubelet HTTP server.
func (c *CMDB) ListenAndServe() {

	klog.InfoS("Starting to listen", "port", c.port)
	s := &http.Server{
		// Addr:           net.JoinHostPort(address.String(), strconv.FormatUint(uint64(port), 10)),
		Addr:           net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(c.port), 10)),
		Handler:        c.server.RestfulCont,
		IdleTimeout:    90 * time.Second, // matches http.DefaultTransport keep-alive timeout
		ReadTimeout:    4 * 60 * time.Minute,
		WriteTimeout:   4 * 60 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	klog.Infof("server http on http://localhost:%d", c.port)
	if c.tls != nil {
		s.TLSConfig = c.tls.Config
		// Passing empty strings as the cert and key files means no
		// cert/keys are specified and GetCertificate in the TLSConfig
		// should be called instead.
		if err := s.ListenAndServeTLS(c.tls.CertFile, c.tls.KeyFile); err != nil {
			klog.ErrorS(err, "Failed to listen and serve")
			os.Exit(1)
		}
	} else if err := s.ListenAndServe(); err != nil {
		klog.ErrorS(err, "Failed to listen and serve")
		os.Exit(1)
	}
}
