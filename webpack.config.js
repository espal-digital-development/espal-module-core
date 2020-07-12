const path = require('path');
const { VueLoaderPlugin } = require('vue-loader');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');

module.exports = (env = {}) => ({
    mode: env.prod ? 'production' : 'development',
    devtool: 'source-map',
    entry: path.resolve(__dirname, './app/assets/vue/main.js'),
    output: {
        path: path.resolve(__dirname, './app/assets/js/spa/'),
        publicPath: '/js/spa/'
    },
    resolve: {
        alias: {
            // this isn't technically needed, since the default `vue` entry for bundlers
            // is a simple `export * from '@vue/runtime-dom`. However having this
            // extra re-export somehow causes webpack to always invalidate the module
            // on the first HMR update and causes the page to reload.
            vue: '@vue/runtime-dom'
        }
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                use: 'vue-loader'
            },
            {
                test: /\.png$/,
                use: {
                    loader: 'url-loader',
                    options: { limit: 8192, name: 'images/[name].[ext]' }
                }
            },
            {
                test: /\.css$/,
                use: [
                    {
                        loader: MiniCssExtractPlugin.loader,
                        options: { hmr: !env.prod }
                    },
                    'css-loader'
                ]
            }
        ]
    },
    plugins: [
        new VueLoaderPlugin(),
        new MiniCssExtractPlugin({
            filename: '[name].css'
        })
    ]
});
