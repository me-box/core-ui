<template>
	<div>
		<div class="mdc-menu-surface--anchor">
			<div class="material-icons mdc-top-app-bar__action-item" @click="openMenu">more_vert</div>
			<div id="app-menu" class="mdc-menu-surface mdc-list">
				<div class="mdc-list-item" v-on:click="viewLogs">
					<span class="mdc-list-item__graphic material-icons" aria-hidden="true">info</span>
					View Logs
				</div>
				<div class="mdc-list-item" v-on:click="restart">
					<span class="mdc-list-item__graphic material-icons" aria-hidden="true">refresh</span>
					Restart
				</div>
				<div class="mdc-list-item" v-on:click="remove">
					<span class="mdc-list-item__graphic material-icons" aria-hidden="true">delete</span>
					Uninstall
				</div>
			</div>
		</div>
	</div>
</template>
<script>
	import {MDCMenuSurface} from '@material/menu-surface';

	export default {
		name: 'appUI',
		props: ['app'],
		data: function () {
			return {
				appMenu: null,
			}
		},
		mounted() {
			this.$parent.setTitle(this.app, true);
			this.appMenu = new MDCMenuSurface(document.querySelector('#app-menu'));
		},
		methods: {
			openMenu: function () {
				this.appMenu.open = true;
			},
			viewLogs: function () {
				alert("Comming soon!!")
			},
			restart: function () {
				fetch(`/core-ui/ui/api/restart`, {
					method: "POST",
					credentials: "same-origin",
					headers: {
						'Accept': 'application/json, text/plain, */*',
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						name: this.app,
					}),
				})
					.then((response) => {
						if (response.status === 401) {
							localStorage.setItem('databoxAuthenticated', false);
							this.$router.push('/');
							return
						}
						if (response.status !== 200) {
							alert("Error restarting app. Sorry.")
						}
						else {
							this.$router.push("MyApps")
						}
					})
					.catch(error => console.error(error));
			},
			remove: function () {
				fetch(`/core-ui/ui/api/uninstall`, {
					method: "POST",
					credentials: "same-origin",
					headers: {
						'Accept': 'application/json, text/plain, */*',
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						name: this.ui,
					}),
				})
					.then((response) => {
						if (response.status === 401) {
							localStorage.setItem('databoxAuthenticated', false);
							this.$router.push('/');
							return
						}
						if (response.status !== 200) {
							alert("Error restarting app. Sorry.")
						}
						else {
							this.$router.push("MyApps")
						}
					})
					.catch(error => console.error(error));
			}
		}
	}
</script>