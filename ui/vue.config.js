// vue.config.js
module.exports = {
	baseUrl: '/core-ui/ui/',
	css: {
		loaderOptions: {
			sass: {
				includePaths: [__dirname  + '/node_modules'],
			},
		},
	},
	pwa: {
		name: 'Databox Dashboard',
		themeColor: '#536878',
		msTileColor: '#536878',
		appleMobileWebAppCapable: 'yes',
		iconPaths: {
			favicon16: 'icons/icon-16.png',
			favicon32: 'icons/icon-32.png',
			appleTouchIcon: 'icons/apple-icon-152.png',
			msTileImage: 'icons/icon-144.png',
			maskIcon: 'icons/safari-pinned-tab.svg'
		}
	}
};