<template>
	<div>
		<div class="mdc-menu-surface--anchor">
			<div class="material-icons mdc-top-app-bar__action-item" @click="openMenu">more_vert</div>
			<div id="app-menu" class="mdc-menu-surface mdc-list">
				<!--<div class="mdc-list-item" v-on:click="viewLogs">
					<span class="mdc-list-item__graphic material-icons" aria-hidden="true">info</span>
					View Logs
				</div>-->
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
				this.$parent.apiRequest('/core-ui/ui/api/restart', {}, {
					method: "POST",
					headers: {
						'Accept': 'application/json, text/plain, */*',
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						name: this.app,
					}),
				})
					.then(() => {
						this.$router.push("/")
					})
			},
			remove: function () {
				this.$parent.apiRequest('/core-ui/ui/api/uninstall', {}, {
					method: "POST",
					headers: {
						'Accept': 'application/json, text/plain, */*',
						'Content-Type': 'application/json'
					},
					body: JSON.stringify({
						name: this.app,
					}),
				})
					.then(() => {
						this.$router.push("/")
					})
			}
		}
	}
</script>