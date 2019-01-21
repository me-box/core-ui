<template>
</template>
<script>
	export default {
		name: 'scanQR',
		mounted: function () {
			this.$parent.title = 'Scan QR Code';
			this.$parent.backRoute = '/login';
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
							console.error(JSON.stringify(err));
						} else {
							body.style.backgroundColor = '';
							QRScanner.destroy();
							const auth = JSON.parse(text);
							let boxIP = auth.ip;
							if (boxIP === "127.0.0.1" || ip.startsWith("169.254.")) {
								for (const ip of auth.ips) {
									if (ip === "127.0.0.1" || ip.startsWith("169.254.")) {
										continue;
									}
									boxIP = ip;
									break;
								}
							}
							let token = auth.token.trim().replace(/^(Token=)/, "");
							this.$router.replace({
								name: 'login',
								params: {url: boxIP, password: token, autoLogin: true}
							});
						}
					});

					QRScanner.show(() => {
						body.style.backgroundColor = 'transparent';
					});
				}
			});
		},
		beforeDestroy: function () {
			let body = document.getElementsByTagName("body")[0];
			body.style.backgroundColor = '';
			QRScanner.destroy();
		}
	}
</script>
<style scoped>
</style>