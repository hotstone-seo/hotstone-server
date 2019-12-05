module.exports = {
  entry: [
    './src/index.js'
  ],
  output: {
    path: __dirname + '/lib',
    filename: "index.js"
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: "babel-loader",
          options: {
            presets: ['@babel/preset-env']
          }
        }
      }
    ]
  },
  // plugins: []
};
