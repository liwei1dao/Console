<template>
    <div>
        <base-v-component
            heading="clusterconsole"
            destring="clusterconsoledes"
        />
         <v-row>
            <v-col
              cols="12"
              sm="12"
              lg="6"
              v-for="(item,index) in clusterservice" :key="index"
            >
              <ServiceCard
                :service="item" 
              >
              </ServiceCard>
            </v-col>
         </v-row>
    </div>
</template>

<script>
import ServiceCard from "./components/MaterialServiceCard"
import { mapGetters } from 'vuex'
export default {
    name:"Services",
    components:{
      ServiceCard
    },
    computed: {
      ...mapGetters([
        'clusterservice'
      ])
    },
    methods: {
      starttimer(){
        if (this.timer == null){
          this.timer = setInterval(() => {
            this.$store.dispatch('console/getconsolecluster')
          }, 1000)
        }

      },
      stoptimer(){
        clearInterval(this.timer);
      }
    },
    created () {
      // this.starttimer()
      this.$store.dispatch('console/getconsolecluster')
    },
    beforeDestroy () {
      //  this.stoptimer()
    }
}
</script>