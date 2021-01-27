<template>
    <v-list dense>
        <template  v-for="item in currRouteGroup">
            <template v-if="item.children && item.children.length > 1">
               
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
            <!-- <v-list-group
                v-if="item.children && item.children.length > 1"
                :key="item.path"
                v-model="item.model"
                :prepend-icon="item.model ? item.meta.icon : item['icon-alt']"
                append-icon=""
            >
                <template v-slot:activator>
                <v-list-item-content>
                    <v-list-item-title>
                    {{ item.meta.title }}
                    </v-list-item-title>
                </v-list-item-content>
                </template>
                <v-list-item
                v-for="(child, i) in item.children"
                :key="i"
                link
                >
                <v-list-item-action v-if="child.icon">
                    <v-icon>{{ child.icon }}</v-icon>
                </v-list-item-action>
                <v-list-item-content>
                    <v-list-item-title>
                    {{ child.text }}
                    </v-list-item-title>
                </v-list-item-content>
                </v-list-item>
            </v-list-group>
            <v-list-item
                v-else-if="item.children && item.children.length == 1"
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
            </v-list-item> -->
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