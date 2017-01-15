module.exports = {
    entry: './src/app/index.js',
    output: {
        path: __dirname + "/public",
        filename: "app.js"
    },
    module: {
        loaders: [
            {
                test: /\.jsx?$/,
                loader: 'babel',
                query: {
                    presets: ['es2015']
                },
                exclude: '/node_modules/'
            }
        ]
    }
};