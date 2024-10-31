// next.config.mjs

/** @type {import('next').NextConfig} */
const nextConfig = {
  webpack: (config) => {
    config.module.rules.push({
      test: /\.svg$/,
      use: [
        {
          loader: '@svgr/webpack',
          options: {
            icon: true, // Optional: makes SVGs easier to use as icons
          },
        },
      ],
    });
    return config;
  },
};

export default nextConfig;
