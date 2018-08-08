<template>
  <div id="Install">
    <div id="wrapper">
        <div v-if="manifest == null || dataSources == null">Loading ..... </div>
        <div v-if="manifest != null && dataSources != null" >
            <h1>Install {{ manifestName }}?</h1>
            <table>
                <tr>
                    <td class="td-icon" rowspan="4"><icon v-bind:name="manifestName" v-bind:displayName="false"/></td>
                    <td><h3>{{ manifest.name }}</h3></td>
                </tr>
                <tr>
                    <td><h3>{{ manifest.author }}</h3></td>
                </tr>
                <tr>
                    <td>{{ manifest.version }}</td>
                </tr>
                <tr>
                    <td>
                        <font-awesome-icon icon="star" />
                        <font-awesome-icon icon="star" />
                        <font-awesome-icon icon="star" />
                        <font-awesome-icon icon="star" />
                        <font-awesome-icon icon="star" class="disabled"/>
                        &nbsp; Moderate Risk
                    </td>
                </tr>
            </table>
            <div><strong>{{ manifest.description }}</strong></div>
        </div>
        <div v-if="manifest != null && manifest.datasources && dataSources != null">
            <h3>Assignee Data Sources to App</h3>
            <datasourcesDDL v-for="ds in manifest.datasources" v-bind:dataSource="ds" v-bind:avalableDataSources="dataSources" v-bind:key="ds.id" v-on:selcted="onDsSelected" />
        </div>
        <div v-if="manifest != null && dataSources != null" >
            <button class="cancel" v-on:click="cancel">Cancel</button>
            <button :class="this.installEnabled + ' install '" v-on:click="install" >Install</button>
        </div>
    </div>
  </div>
</template>
<script>

import Icon from '../components/renderAppDriverIcon.vue'
import datasourcesDDL from '../components/renderDataSourcesDDL.vue'

import testdata from '../testData/DataSources.json'
import testManifest from '../testData/app-manifest.json'

export default {
  name: 'install',
  props: {
   },
  components: {Icon, datasourcesDDL},
  data() {
    return {
        manifestName: "?",
        selectedDataSources: {},
        installEnabled: "disabled",
        loadingMan: "true",
        loadingData: "true"
    }
  },
  asyncComputed: {
    manifest: {
        get() {

            if (this.manifestName == "?") {
                return Promise.resolve(null)
            }

            return fetch('/core-ui/ui/api/manifest/' + this.manifestName, {credentials: "same-origin"})
            .then((response) => {
                if (response.status == 401) {
                    localStorage.setItem('databoxAuthenticated', false)
                    this.$router.push('/')
                    return
                }
                return response.json()
            })
            .then((json) =>  {
                return json;
            })
            .catch(() => {
                return new Promise((resolve,reject) => { setTimeout(()=>{ resolve(testManifest) },500) });
            });
        },
        default: null,
    },
    dataSources: {
        get() {
            return fetch('/core-ui/ui/api/dataSources', {credentials: "same-origin"})
            .then((response) => {
                if (response.status == 401) {
                    localStorage.setItem('databoxAuthenticated', false)
                    this.$router.push('/')
                    return
                }
                return response.json()
            })
            .then((json) =>  {
                return json;
            })
            .catch(() => {
                return new Promise((resolve,reject) => { setTimeout(()=>{ resolve(testdata) },500) });
            });
        },
        default: null,
    },
  },
  updated () {
    this.updateInstallButtonState()
  },
  mounted() {

    //cache context
    let _this = this

    //get manifest name from Query string
    this.manifestName = this.$route.query['manifest']

    this.updateInstallButtonState()

  },
  methods: {
      onDsSelected: function (ds,Clientid) {
          if (ds.dsType) {
            this.selectedDataSources[Clientid] = ds
          } else {
            this.selectedDataSources[Clientid] = undefined
            delete this.selectedDataSources[Clientid]
          }
          this.updateInstallButtonState()
      },
      install : function () {
          if (this.installEnabled == "enabled") {
            if(Array.isArray(this.manifest.datasources)) {
                this.manifest.datasources.forEach((ds,index,arr)=>{
                    arr[index].hypercat = this.selectedDataSources[ds["clientid"]].hypercat
                });
            }
            let data = {
                manifest: this.manifest,
            }
            fetch(`/core-ui/ui/api/install`, {
                    method: "POST",
                    credentials: "same-origin",
                    headers: {
                        'Accept': 'application/json, text/plain, */*',
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(data),
                })
                .then( (response) => {
                    if (response.status == 401) {
                        localStorage.setItem('databoxAuthenticated', false)
                        this.$router.push('/')
                        return
                    }
                    if (response.status != 200) {
                        alert("Error installing app. Sorry.")
                    } else {
                        this.$router.push("MyApps")
                    }
                })
                .catch(error => console.error(error));
            return
          }

          //else
          alert("Please select the required data sources")

      },
      cancel: function (){
          this.$router.push("AppStore")
      },
      updateInstallButtonState: function() {

          if(this.manifest == null) {
            this.installEnabled = 'disabled'
            return
          }

          //TODO deal with optional datasources
          //TODO Add functionality for min and max datasources of a given type
          if (!this.manifest.datasources || Object.keys(this.selectedDataSources).length == this.manifest.datasources.length) {
            this.installEnabled = 'enabled'
          } else {
            this.installEnabled = 'disabled'
          }
      }
  }
}
</script>
<style>
table {
    table-layout: fixed;
    border-collapse: collapse;
    width: 100%;
}
td {
    overflow: hidden;
    white-space: nowrap;
    padding-left: 5px;
}
.td-icon {
    width:150px;
}
.disabled {
    color: #909090;
}
div {
  padding: 5px;
  display: block;
}
.cancel {
	-moz-box-shadow:inset 0px 39px 0px -24px #e67a73;
	-webkit-box-shadow:inset 0px 39px 0px -24px #e67a73;
	box-shadow:inset 0px 39px 0px -24px #e67a73;
	background-color:#e4685d;
	-moz-border-radius:4px;
	-webkit-border-radius:4px;
	border-radius:4px;
	border:1px solid #ffffff;
	display:inline-block;
	cursor:pointer;
	color:#ffffff;
	font-family:Arial;
	font-size:15px;
	padding:6px 15px;
	text-decoration:none;
	text-shadow:0px 1px 0px #b23e35;
}
.install {
    -moz-box-shadow:inset 0px 39px 0px -24px #77b55a;
	-webkit-box-shadow:inset 0px 39px 0px -24px #77b55a;
	box-shadow:inset 0px 39px 0px -24px #77b55a;
	background-color:#77b55a;
	-moz-border-radius:4px;
	-webkit-border-radius:4px;
	border-radius:4px;
	border:1px solid #ffffff;
	display:inline-block;
	cursor:pointer;
	color:#ffffff;
	font-family:Arial;
	font-size:15px;
	padding:6px 15px;
	text-decoration:none;
	text-shadow:0px 1px 0px #77b55a;
}
button {
    margin: 5px;
    min-width: 100px;
}
button:active {
	position:relative;
	top:1px;
}

button.disabled {
     -moz-box-shadow:inset 0px 39px 0px -24px #808080;
	-webkit-box-shadow:inset 0px 39px 0px -24px #808080;
	box-shadow:inset 0px 39px 0px -24px #808080;
	background-color:#808080;
  	text-shadow:0px 1px 0px #808080;
}
</style>
