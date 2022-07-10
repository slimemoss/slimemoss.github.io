const webpack = require('webpack');
var path = require('path');

module.exports = {
    entry: {
        index: './src/index.tsx',
    },
    output: {
        path: path.join(__dirname, 'dist'),
        filename: '[name].js',
    },
    module: {
        rules: [
            {
                test: /\.jsx?$/,
                use: ['babel-loader'],
                exclude: /node_modules/,
            },
            {
                test: /\.tsx?$/,
                use: ['ts-loader'],
                exclude: /node_modules/,
            },
            {
                test: /\.css?$/,
                use: ['style-loader', 'css-loader'],
            },
            {
              test: /\.(csv)$/i,
              use: [
                {
                  loader: 'raw-loader'
                }
              ],
            },            {
              test: /\.(png|jpe?g|gif)$/i,
              use: [
                {
                  loader: 'file-loader',
                  options: {
                    name: '[path][name].[ext]'
                  }
                }
              ],
            }
        ]
    },
    resolve: {
        extensions: ['.js', '.jsx', '.ts', '.tsx', '.json'],
        modules: ['node_modules']
    },
    mode: 'development',
    devServer: {
        port: 8080,
        host: '0.0.0.0',
        historyApiFallback: true,
        static: __dirname,
        hot: true,
        liveReload: true,
    },
    watchOptions: {
        aggregateTimeout: 200,
        poll: 1000,
        ignored: '**/node_modules'
    },
    performance: {
        maxEntrypointSize: 1000000,
        maxAssetSize: 1000000
    }
};
