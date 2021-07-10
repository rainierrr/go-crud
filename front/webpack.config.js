const path = require('path');

module.exports = {
    entry: path.resolve(__dirname, "./src/index.tsx"),
    output: {
        path: path.resolve(__dirname, "dist"),
        filename: "bundle.js"
    },
    resolve: {
        extensions: [".ts", ".tsx", ".js", ".jsx", ".json"]
    },
    devServer: {
        contentBase: path.join(__dirname, 'dist'),
        host: '0.0.0.0',
        inline: true,
        port: 3000,
        public: 'localhost:3030'
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                loader:'ts-loader'
            }
        ]
    },
    stats: {
        errorDetails: true
    }
}