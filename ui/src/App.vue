<template>
	<div id="app">
		<header class="mdc-top-app-bar" v-if="authenticated">
			<div class="mdc-top-app-bar__row">
				<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
					<a id="toolbarBack" v-on:click="back" class="material-icons mdc-top-app-bar__navigation-icon">
						arrow_back
					</a>
					<span id="toolbarTitle" class="mdc-top-app-bar__title">Databox Dashboard</span>
				</section>
				<section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end">
					<router-view name="toolbar" style="display: flex"></router-view>
					<div class="material-icons mdc-top-app-bar__action-item" v-if="notifications.length === 0">
						notifications_none
					</div>
					<div class="mdc-menu-surface--anchor">
						<div class="material-icons mdc-top-app-bar__action-item" @click="openNotifications"
						     v-if="notifications.length > 0">
							notifications
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
			<div id="content">
				<router-view></router-view>
			</div>
		</div>
	</div>
</template>

<script>
	import {MDCMenuSurface} from '@material/menu-surface';

	export default {
		name: 'app',
		data() {
			return {
				databoxUrl: "localhost",
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
					this.authenticated = true;
				})
		},
		mounted() {
			this.authenticated = localStorage.getItem("databoxAuthenticated") === "true";
			if (this.authenticated) {
				this.notificationMenu = new MDCMenuSurface(document.querySelector('#notification-menu'));
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
						if (err.status === 401 || err.status === 404) {
							this.authenticated = false;
							this.$router.replace("/login");
						} else if (this.isDev) {
							return cannedData
						} else {
							throw err;
						}
					});
			},
			logout() {
				this.authenticated = false;
				localStorage.setItem('databoxAuthenticated', 'false');
				this.$router.replace("/");
			},
			login(url, password) {
				this.databoxUrl = url;
				localStorage.setItem('databoxURL', url);
				fetch('https://' + url + '/core-ui/ui/api/connect', {
					method: "GET",
					credentials: "include",
					mode: "cors",
					headers: {
						'Authorization': "Token " + password,
					},
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
					.catch(() => {
						if (this.isDev) {
							this.authenticated = true;
							this.$router.replace("/")
						} else {
							this.authenticated = false;
						}
					});
			},
			goto(page) {
				this.$router.push(page);
			},
			back() {
				this.$router.go(-1);
			},
			openNotifications() {
				console.log(this.notificationMenu);
				this.notificationMenu.open = true;
			},
			setTitle(title, backHidden) {
				let titleBar = document.getElementById('toolbarTitle');
				if (titleBar != null) {
					titleBar.innerText = title;
				}
				document.title = title;
				let toolbarBack = document.getElementById('toolbarBack');
				if (toolbarBack !== null) {
					if (backHidden) {
						toolbarBack.style.visibility = 'hidden';
					}
					else {
						toolbarBack.style.visibility = 'visible';
					}
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
