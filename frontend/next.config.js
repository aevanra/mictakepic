module.exports = {
  webpack: (config, { isServer }) => {
    // Allow using require.context in client-side code
    if (!isServer) {
      config.module.rules.push({
        test: /\*.(png|jpg|gif|svg)$/i,
        use: [
          {
            loader: 'file-loader',
            options: {
              publicPath: '/_next/static/images',
              outputPath: 'static/images',
              name: '[name].[hash].[ext]',
            },
          },
        ],
      });
    }
    return config;
  },
};
