const path = require('path');
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
    contentBase: path.join(__dirname,'dist')
  },
  module: {
    rules: [
      {
        test:/\.ts$/,loader:'ts-loader'
      }
    ]
  }
};
