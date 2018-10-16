// vue.config.js
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