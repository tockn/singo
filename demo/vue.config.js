module.exports = {
  transpileDependencies: ["vuetify"],
  configureWebpack: {
    devServer: {
      watchOptions: {
        poll: true
      }
    }
  },
  chainWebpack: config => config.resolve.symlinks(false)
};
