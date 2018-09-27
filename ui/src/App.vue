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
							<div class="mdc-list-item" v-for="item in notifications">
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
		data: function () {
			return {
				authToken: null,
				authenticated: false,
				notifications: ["A Notification", "Another Notification"],
				notificationMenu: null,
			}
		},
		created: function () {
			let authenticated = localStorage.getItem('databoxAuthenticated');
			if (authenticated === "true") {
				this.authenticated = true;
			}
			else {
				this.$router.push('/login');
			}
		},
		mounted: function () {
			this.notificationMenu = new MDCMenuSurface(document.querySelector('#notification-menu'));
		},
		methods: {
			goto: function (page) {
				this.$router.push(page);
			},
			back: function () {
				this.$router.go(-1);
			},
			openNotifications: function () {
				console.log(this.notificationMenu);
				this.notificationMenu.open = true;
			},
			setTitle: function (title, backHidden) {
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

	@import "@material/button/mdc-button";
	@import "@material/card/mdc-card";
	@import "@material/dialog/mdc-dialog";
	@import "@material/list/mdc-list";
	@import "@material/menu-surface/mdc-menu-surface";
	@import "@material/select/mdc-select";
	@import "@material/textfield/mdc-text-field";
	@import "@material/theme/mdc-theme";
	@import "@material/top-app-bar/mdc-top-app-bar";

	html {
		height: 100%;
	}

	body {
		height: 100%;
		margin: 0;
	}

	#content {
		display: flex;
		justify-content: center;
		padding: 24px;
		flex-direction: column;
		align-items: center;
	}

	@media (max-width: 479px) {
		#content {
			padding: 16px;
			align-items: stretch;
		}
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
