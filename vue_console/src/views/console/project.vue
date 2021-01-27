<template>
    <div v-if="projectinfo">
        <section class="mb-12 text-center">
            <h1
            class="font-weight-light mb-2 headline"
            v-text="projectinfo.ProjectName"
            />

            <span
            class="font-weight-light subtitle-1"
            >
            {{projectinfo.ProjectDes}}
            </span>
        </section>
        <base-v-divider size="small" orientation="left">{{$t('projectmember')}}</base-v-divider>
        <v-simple-table>
            <template v-slot:default>
            <thead>
                <tr>
                <th class="text-left">
                    {{$t('projectposition')}}
                </th>
                <th class="text-left">
                    {{$t('projectname')}}
                </th>
                </tr>
            </thead>
            <tbody>
                <tr
                v-for="(member,key) in projectinfo.ProjectMember"
                :key="key"
                >
                <td>{{ member.ProjectPosition }}</td>
                <td>{{ member.MemberName }}</td>
                </tr>
            </tbody>
            </template>
        </v-simple-table>

        <p class="text--disabled text-right">{{$t('version')+':'+projectinfo.ProjectVersion}}</p> 
        <p class="text--disabled text-right">{{projectinfo.ProjectTime}}</p> 
    </div>
</template>

<script>
// import HostCard from "./components/MaterialHostCard"
// import PieChart from "@/components/PieChart"
// import LineChart from "@/components/LineChart"
import moment from 'moment'
import { mapGetters } from 'vuex'
export default {
    name:"Base",
    components:{
        // HostCard,
        // PieChart,
        // LineChart
    },
    computed: {
      ...mapGetters([
        'projectinfo',
      ]),

    },
    data :function(){
      return {
        tab:null,
      }
    },
    filters: {
        formatDate(time) {
            var date = new Date(time*1000);
            return moment(date).format("YYYY-MM-DD HH:mm:ss");
        }
    },
    methods: {
      starttimer(){
        if (this.timer == null){
          this.timer = setInterval(() => {
            this.$store.dispatch('console/gethostinfo', this.form)
          }, 1000)
        }
      },
      stoptimer(){
        clearInterval(this.timer);
      }
    },
    created () {
      // this.starttimer()
      this.$store.dispatch('console/getprojectinfo')
      this.$store.dispatch('console/gethostinfo')
      this.$store.dispatch('console/getcpuinfo')
      this.$store.dispatch('console/getmemoryinfo')
    },
    beforeDestroy () {
      //  this.stoptimer()
    }
}
</script>

<style lang="sass">

</style>
