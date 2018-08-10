<template>
  <div>
    <h1>Databox Settings</h1>
    <p>Change password ... etc</p>
  </div>
</template>
<script>

export default {
  name: 'settings',
  props: {},
  created: function () {
    let devmode = localStorage.getItem('dev')
    if (this.$parent.authenticated == "false" && devmode != "true") {
      this.$router.push('login')
    }
  }
}

</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

</style>