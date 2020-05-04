module.exports = {
  transpileDependencies: ["vuetify"],
  configureWebpack: {
    devServer: {
      watchOptions: {
        poll: true
      }
    }
  }
};
