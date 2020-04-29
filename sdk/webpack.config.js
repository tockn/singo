const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
module.exports = {
  entry: {
    bundle: './src/app.ts'
  },
  output: {
    path: path.join(__dirname,'dist'),
    filename: '[name].js'  // [name]はentryで記述した名前(この例ではbundle）が入る
  },
  resolve: {
    extensions:['.ts','.js']
  },
  devServer: {
    contentBase: path.join(__dirname,'dist'),
    open: true
  },
  module: {
    rules: [
      {
        test:/\.ts$/,loader:'ts-loader'
      }
    ]
  },
  plugins: [
    new HtmlWebpackPlugin()
  ]
};
