<template>
</template>
<script>
	export default {
		name: 'scanQR',
		props: {},
		data: function () {
			return {
				password: "",
				url: "",
				error: ""
			}
		},
		computed: {
			valid: function () {
				return this.url && this.password;
			}
		},
		mounted: function () {
			QRScanner.prepare((err, status) => {
				let body = document.getElementsByTagName("body")[0];
				if (err) {
					console.error(err);
					body.style.backgroundColor = '';
					QRScanner.destroy();
					this.$router.push('/scan');
				}
				if (status.authorized) {
					QRScanner.scan((err, text) => {
						if (err) {
							console.error(err);
						} else {
							body.style.backgroundColor = '';
							QRScanner.destroy();
							const auth = JSON.parse(text);
							this.$parent.login(auth.ip, auth.token);
						}
					});

					QRScanner.show(function(status){
						console.log(status);
						body.style.backgroundColor = 'transparent';
					});
				}
			});
		},
		beforeDestroy: function() {
			let body = document.getElementsByTagName("body")[0];
			body.style.backgroundColor = '';
			QRScanner.destroy();
		},
		methods: {
			scan: function () {
				console.log("Scan Test 1");
				this.$parent.scan();
			},
			login: function () {
				this.$parent.login(this.url, this.password);
			}
		}
	}
</script>
<style scoped>
</style>