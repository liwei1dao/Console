<template>
    <div>
        <v-tabs v-model="tabchange" background-color="transparent" centered icons-and-text>
            <v-tab class="text-capitalize" v-for="(value,name) in routes" :key="name">
                {{$t(value.title)}}
                <v-icon>{{value.icon}}</v-icon>
            </v-tab>
        </v-tabs>
    </div>
</template>
<script>
import path from 'path'
import { mapGetters } from 'vuex'
export default {
    name : "TopNavigation",
    computed: {
        ...mapGetters([
            'routes',
            'currtagroute',
        ]),
    },
    watch: {
        tabchange: function (index) {
            var keys = Object.keys(this.routes)
            var groute = this.routes[keys[index]]
            var currtagroute = groute.routes[groute.default]
            console.log("keys:"+keys[index]+"|currtagroute.children[0].path:"+currtagroute.children[0].path)
            this.$store.dispatch('router/setcurrtagroute',currtagroute)
            this.$router.push({path:this.resolvePath(currtagroute.path,currtagroute.children[0].path)})
        },
    },
    data() {
        return {
            tabchange:0
        }
    },
    methods:{
        resolvePath(cpath,vpath) {
            return path.resolve(cpath, vpath)
        }
    },

}
</script>