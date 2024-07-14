/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "standalone",
  async rewrites() {
    return [
      {
        source: '/api/:slug*',
        destination: 'http://localhost:18080/:slug*'
      }
    ]
  },
};

export default nextConfig;
