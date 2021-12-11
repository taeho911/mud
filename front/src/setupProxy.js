const { createProxyMiddleware } = require('http-proxy-middleware')

const API_DOMAIN = process.env.API_DOMAIN || 'localhost'
const API_PORT = process.env.API_PORT || '8080'

module.exports = function(app) {
  app.use(
    '/api',
    createProxyMiddleware({
      target: `http://${API_DOMAIN}:${API_PORT}`,
      changeOrigin: true
    })
  )
}
