<template>
  <div>
  <section style="padding:16px">
      <h1>Welcome to Databox</h1>
      <div>Databox lets you take control of your personal data and IoT devices.</div>
      <div>Get started by:</div>

    <p>Install a what you need from the&nbsp;<router-link :to="{ path:'appstore'}">App store</router-link>.<p>

       <div style="display: flex;justify-content: center;align-items: center;"><a href="https://play.google.com/store/apps/details?id=io.databox.app" target="_top" style="margin-top: 5px; margin-right: 20px"><img alt="Get Databox on Google Play" height="60" src="https://play.google.com/intl/en_us/badges/images/generic/en_badge_web_generic.png"></a><a href="https://itunes.apple.com/us/app/databox-controller/id1295071825?mt=8" target="_top" style="display:inline-block;overflow:hidden;background:url(//linkmaker.itunes.apple.com/assets/shared/badges/en-us/appstore-lrg.svg) no-repeat;width:135px;height:40px;background-size:contain;" title="Get Databox on Apple App Store"></a></div><div style="display: flex;justify-content: center;align-items: center;"><a class="mdc-button mdc-button--compact" href="#!/qrcode">APP QR Config Code</a></div>
       <div>For more help see the&nbsp;<a href="https://github.com/me-box/databox">documentation</a>.</div>
    </section>
    </div>
</template>
<script>

export default {
  name: 'home',
  props: {},
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
div {
  padding: 5px;
}
section {
  display: block;
}
</style>