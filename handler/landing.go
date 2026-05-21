package handler

import (
	"net/http"
	"strconv"
)

var landingPageHTML = []byte(`<!doctype html><html lang="en"><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>web.nacl.one</title><style>body{margin:0;min-height:100vh;display:grid;place-items:center;font:16px/1.5 system-ui,-apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif;background:#f6f1e8;color:#1d1a16}main{max-width:44rem;padding:2rem}h1{margin:0 0 .5rem;font-size:2rem;letter-spacing:.04em}p{margin:.4rem 0}code,pre{font:inherit;background:#ece4d8;padding:.1rem .3rem;border-radius:.25rem}.grid{display:grid;grid-template-columns:repeat(3,minmax(0,1fr));gap:.75rem;margin-top:1.25rem}.card{padding:.8rem 1rem;background:#fff;border:1px solid #ded2c2;border-radius:.75rem}.atom{display:flex;justify-content:center;align-items:center;min-height:6rem;font-size:2.1rem;font-weight:700;letter-spacing:.12em}.atom span{position:relative}.atom span:before,.atom span:after{content:"";position:absolute;inset:-1.6rem;border:1px solid #cdb89f;border-radius:999px;transform:scaleX(1.2)}.atom span:after{inset:-2.4rem}.small{font-size:.92rem;color:#5a5247}@media (max-width:700px){.grid{grid-template-columns:1fr}.atom{min-height:5rem;font-size:1.8rem}}</style><main><h1>NACL</h1><p>web.nacl.one is a lightweight HTTPS proxy endpoint. When opened directly in a browser, it shows a small note about salt: <code>NaCl</code>, sodium chloride.</p><div class="grid"><section class="card"><div class="atom"><span>NaCl</span></div><p class="small">An ionic crystal lattice of sodium and chloride ions. Clean, stable, and familiar.</p></section><section class="card"><strong>History</strong><p class="small">Salt shaped trade, preservation, and taxation long before modern chemistry gave it a formula.</p></section><section class="card"><strong>Properties</strong><p class="small">White crystalline solid, soluble in water, and essential in labs, kitchens, and industry.</p></section></div><p class="small">For proxy use, configure your client to speak HTTPS proxy to this host.</p></main>`)

func isLandingPageRequest(req *http.Request) bool {
	if req == nil || req.URL == nil {
		return false
	}
	if req.Method != http.MethodGet && req.Method != http.MethodHead {
		return false
	}
	if req.URL.Scheme != "" || req.URL.Host != "" {
		return false
	}
	return req.Host != ""
}

func serveLandingPage(wr http.ResponseWriter, req *http.Request) {
	wr.Header().Set("Content-Type", "text/html; charset=utf-8")
	wr.Header().Set("Cache-Control", "no-store")
	wr.Header().Set("X-Content-Type-Options", "nosniff")
	wr.Header().Set("Content-Length", strconv.Itoa(len(landingPageHTML)))
	wr.WriteHeader(http.StatusOK)
	if req != nil && req.Method == http.MethodHead {
		return
	}
	_, _ = wr.Write(landingPageHTML)
}
