<template>
	<div id="app">
		<header class="mdc-top-app-bar" v-if="authenticated">
			<div class="mdc-top-app-bar__row">
				<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
					<a v-if="backRoute" v-on:click="back" class="material-icons mdc-top-app-bar__navigation-icon">
						arrow_back
					</a>
					<span class="mdc-top-app-bar__title">{{ title }}</span>
				</section>
				<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end">
					<router-view name="toolbar" style="display: flex"></router-view>
					<div class="mdc-menu-surface--anchor">
						<div class="material-icons mdc-top-app-bar__action-item" @click="openNotifications">
							{{ (this.notifications.length > 0 ) ? 'notifications' : 'notifications_none'}}
						</div>

						<div id="notification-menu" class="mdc-menu-surface mdc-list mdc-list--two-line">
							<div class="mdc-list-item" v-for="item in notifications" :key="item">
								<div class="mdc-list-item__text">
									<div class="mdc-list-item__primary-text">{{item}}</div>
									<div class="mdc-list-item__secondary-text">Yesterday, 1:30pm</div>
								</div>
							</div>
						</div>
					</div>
				</section>
			</div>
		</header>

		<div class="mdc-top-app-bar--fixed-adjust">
			<div v-if="connecting" id="content">
				<Spinner />
			</div>
			<div v-else id="content">
				<router-view />
			</div>
		</div>
	</div>
</template>

<script>
	import {MDCMenuSurface} from '@material/menu-surface';
	import Spinner from "./components/Spinner";

	export default {
		name: 'app',
		components: {Spinner},
		data() {
			return {
				title: 'Databox Dashboard',
				backRoute: null,
				databoxUrl: "localhost",
				connecting: true,
				authenticated: false,
				notifications: [],
				notificationMenu: null,
			}
		},
		created() {
			if (localStorage.getItem("databoxUrl") !== null) {
				this.databoxUrl = localStorage.getItem("databoxUrl");
			} else if (!this.isMobile) {
				this.databoxUrl = window.location.hostname;
			}
			this.apiRequest('/core-ui/ui/api/containerStatus', {})
				.then(() => {
					this.connecting = false;
					this.authenticated = true;
				})
		},
		updated() {
			if (this.authenticated && this.notificationMenu == null) {
				this.notificationMenu = new MDCMenuSurface(document.getElementById('notification-menu'));
			} else {
				this.notificationMenu = null;
			}
		},
		watch: {
			title(val) {
				document.title = val
			}
		},
		methods: {
			apiRequest: function (url, cannedData, opts = {}) {
				opts.credentials = 'include';
				opts.mode = 'cors';
				return fetch('https://' + this.databoxUrl + url, opts)
					.then((response) => {
						if (!response.ok) {
							throw {message: response.statusText, status: response.status};
						} else {
							return response;
						}
					})
					.then((response) => {
						return response.json()
					})
					.catch((err) => {
						this.connecting = false;
						if (err.status === 401 || err.status === 404) {
							this.authenticated = false;
							this.$router.replace("/login");
							throw err;
						} else if (this.isDev) {
							return cannedData
						} else {
							throw err;
						}
					});
			},
			logout() {
				this.authenticated = false;
				this.$router.replace("/login");
			},
			login(url, password) {
				this.databoxUrl = url;
				localStorage.setItem('databoxURL', url);
				return fetch('https://' + url + '/core-ui/ui/api/connect', {
					method: "GET",
					credentials: "include",
					mode: "cors",
					headers: {
						'Authorization': "Token " + password,
					},
				})
					.then((response) => {
						if (!response.ok) {
							throw {message: response.statusText, status: response.status};
						} else {
							return response;
						}
					})
					.then((response) => {
						return response.text()
					})
					.then((body) => {
						if (body === "connected") {
							this.authenticated = true;
							localStorage.setItem('databoxAuthenticated', 'true');
							this.$router.replace("/")
						} else {
							return Promise.reject("Auth failed")
						}
					})
					.catch((error) => {
						if (this.isDev) {
							this.authenticated = true;
							this.$router.replace("/")
						} else {
							this.authenticated = false;
							throw error;
						}
					});
			},
			back() {
				this.$router.push(this.backRoute);
			},
			openNotifications() {
				if(this.notificationMenu != null && this.notifications.length > 0) {
					this.notificationMenu.open = true;
				}
			}
		}
	}
</script>

<style lang="scss">
	$mdc-theme-primary: #536878;
	$mdc-theme-accent: #41b883;
	$mdc-theme-background: #fff;

	@import "~@material/button/mdc-button";
	@import "~@material/card/mdc-card";
	@import "~@material/dialog/mdc-dialog";
	@import "~@material/list/mdc-list";
	@import "~@material/menu-surface/mdc-menu-surface";
	@import "~@material/select/mdc-select";
	@import "~@material/textfield/mdc-text-field";
	@import "~@material/theme/mdc-theme";
	@import "~@material/top-app-bar/mdc-top-app-bar";

	html {
		height: 100%;
	}

	body {
		height: 100%;
		margin: 0;
	}

	#app {
		height: 100%;
	}

	#content {
		height: 100%;
		display: flex;
		justify-content: flex-start;
		flex-direction: column;
		align-items: center;
	}

	@media (max-width: 479px) {
		#content {
			align-items: stretch;
		}
	}

	.mdc-top-app-bar--fixed-adjust {
		box-sizing: border-box;
		height: 100%;
	}

	#toolbarTitle {
		font-family: 'Avenir', Helvetica, Arial, sans-serif;
	}

	#app {
		font-family: 'Avenir', Helvetica, Arial, sans-serif;
		-webkit-font-smoothing: antialiased;
		-moz-osx-font-smoothing: grayscale;
		text-align: left;
		color: #2c3e50;
		width: 100%;
	}
</style>
