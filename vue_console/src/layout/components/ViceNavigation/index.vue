<template>
    <v-list dense>
        <template  v-for="item in currRouteGroup">
            <template v-if="item.children && item.children.length > 1">
               <v-list-group :key="item.meta.title">
                    <template v-slot:activator>
                        <v-list-item-action>
                            <v-icon>{{item.meta.icon}}</v-icon>
                        </v-list-item-action>
                        <v-list-item-content>
                            <v-list-item-title v-text="$t(item.meta.title)"></v-list-item-title>
                        </v-list-item-content>
                    </template>
                    <v-list-item
                        v-for="child in item.children"
                        :key="child.meta.title"
                        :to="resolvePath(child.path)"
                        link
                    >
                        <v-list-item-action>
                        </v-list-item-action>
                        <v-list-item-title v-text="$t(child.meta.title)"></v-list-item-title>
                        <v-list-item-icon>
                            <v-icon v-text="child.meta.icon"></v-icon>
                        </v-list-item-icon>
                    </v-list-item>
               </v-list-group>
            </template>
            <template v-else-if="item.children && item.children.length == 1">
                <v-list-item
                    :key="item.children[0].path"
                    :to="resolvePath(item.children[0].path)"
                    link
                >
                    <v-list-item-action>
                    <v-icon>{{item.children[0].meta.icon}}</v-icon>
                    </v-list-item-action>
                    <v-list-item-content>
                    <v-list-item-title>
                        {{$t(item.children[0].meta.title)}}
                    </v-list-item-title>
                    </v-list-item-content>
                </v-list-item>
            </template>
            <template v-else>

            </template>
        </template>
    </v-list>
</template>
<script>
import path from 'path'
import { mapGetters } from 'vuex'
export default {
    name:"ViceNavigation",
    computed: {
        ...mapGetters([
            'routes',
            'currtagroute',
        ]),
        currRouteGroup(){
            return this.routes[this.currtagroute.navigation.name].routes
        }
    },
    data:function(){
        return {

        }
    },
    watch: {
        // $route(route) {
        //     if (route.path.startsWith('/base')) {
        //         return
        //     }
        //     this.getBreadcrumb()
        // }
    },
    methods:{
        resolvePath(vpath) {
            return path.resolve(this.currtagroute.path, vpath)
        }
    }
}
</script>