// vue.config.js
console.log(__dirname, '/node_modules');
module.exports = {
	baseUrl: '',
	css: {
		loaderOptions: {
			sass: {
				includePaths: [__dirname  + '/node_modules'],
			},
		},
	},
};