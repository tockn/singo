const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const environment = process.env.NODE_ENV || 'dev';
module.exports = {
  entry: {
    bundle: './src/index.ts'
  },
  output: {
    path: path.join(__dirname,'dist'),
    filename: 'index.js',
    library: 'singo-sdk',
    libraryTarget: 'umd',
    globalObject: 'typeof self !== \'undefined\' ? self : this'
  },
  resolve: {
    extensions:['.ts','.js'],
    alias: {
      userEnv$: path.resolve(__dirname, `.env/${environment}.ts`),
    }
  },
  devServer: {
    contentBase: path.join(__dirname,'dist'),
    open: true,
    host: 'localhost',
    disableHostCheck: true
  },
  module: {
    rules: [
      {
        test:/\.ts$/,loader:'ts-loader'
      }
    ]
  },
  plugins: [
    // new HtmlWebpackPlugin({
    //   template: './index.html'
    // })
  ],
  optimization: {
    minimize: false,
  }
};
