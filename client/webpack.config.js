

const path = require('path')

module.exports = {
    // モード値を production に設定すると最適化された状態で、
    // development に設定するとソースマップ有効でJSファイルが出力される
    mode: "development",

    // メインとなるJavaScriptファイル（エントリーポイント）
    entry: ['@babel/polyfill', './index.js'],
    output: {
        path: path.resolve(__dirname, './public/'),
        filename: "bundle.js"
    },
    devServer: {
        contentBase: path.join(__dirname, 'public'),
        port: 3500,
        compress: true
    },
    module: {
        rules: [
            {
                // 拡張子 .js の場合
                test: /\.js$/,
                use: [
                    {
                        // Babel を利用する
                        loader: "babel-loader",
                        // Babel のオプションを指定する
                        options: {
                            presets: [
                                // プリセットを指定することで、ES2019 を ES5 に変換
                                "@babel/preset-env",
                                "@babel/react"
                            ]
                        }
                    }
                ]
            }
        ]
    }
};