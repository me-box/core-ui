// vue.config.js
module.exports = {
	publicPath: process.env.VUE_APP_BASE_URL,
	css: {
		loaderOptions: {
			sass: {
				sassOptions: {
					includePaths: [__dirname + '/node_modules']
				}
			}
		}
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