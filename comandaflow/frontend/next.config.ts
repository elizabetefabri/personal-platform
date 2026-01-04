/** @type {import('next').NextConfig} */

const nextConfig = {
  output: "export", 
  images: {
    unoptimized: true, 
  },
  trailingSlash: true, 

  async redirects() {
    return [
      {
        source: '/:path*',
        has: [{ type: 'host', value: 'elizabetesousafabri.com.br' }],
        destination: 'https://elizabetesousafabri.com.br/:path*',
        permanent: true, 
      },
    ];
  },
  async headers() {
    return [
      {
        source: '/(.*)',
        headers: [
          { key: 'Strict-Transport-Security', value: 'max-age=31536000; includeSubDomains; preload' },
        ],
      },
    ];
  },
};

module.exports = nextConfig;
