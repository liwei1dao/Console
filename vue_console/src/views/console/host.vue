<template>
    <div>
      <base-v-component
          heading="hostconsole"
          destring="hostconsoledes"
      />
        <v-col
          cols="12"
        >
          <HostCard
            v-if="hostinfo"
            color="info"
            icon="mdi-hamburger"
            :title="$t('systemdes')"
            :value="$t('system')"
            sub-icon="mdi-clock"
            sub-text="Just Updated"
          >
            <template v-slot:service>
                <v-simple-table>
                  <template v-slot:default>
                  <tbody>
                      <tr><td>{{$t('hostid')}}</td><td>{{hostinfo.HostID}}</td></tr>
                      <tr><td>{{$t('hostname')}}</td><td>{{hostinfo.Hostname}}</td></tr>
                      <tr><td>{{$t('hostuptime')}}</td><td>{{hostinfo.Uptime | secondToDate}}</td></tr>
                      <tr><td>{{$t('hostboottime')}}</td><td>{{hostinfo.BootTime | formatDate}}</td></tr>
                      <tr><td>{{$t('hostprocsnum')}}</td><td>{{hostinfo.Procs}}</td></tr>
                      <tr><td>{{$t('hostos')}}</td><td>{{hostinfo.OS}}</td></tr>
                      <tr><td>{{$t('hostplatform')}}</td><td>{{hostinfo.Platform}}</td></tr>
                      <tr><td>{{$t('hostplatformfamily')}}</td><td>{{hostinfo.PlatformFamily}}</td></tr>
                      <tr><td>{{$t('hostplatformversion')}}</td><td>{{hostinfo.PlatformVersion}}</td></tr>
                      <tr><td>{{$t('hostplatformkernelarch')}}</td><td>{{hostinfo.KernelArch}}</td></tr>
                      <tr><td>{{$t('hostVirtualizationsystem')}}</td><td>{{hostinfo.VirtualizationSystem}}</td></tr>
                      <tr><td>{{$t('hostVirtualizationrole')}}</td><td>{{hostinfo.VirtualizationRole}}</td></tr>
                  </tbody>
                  </template>
              </v-simple-table>
            </template>
          </HostCard>
        </v-col>

        <v-col
          cols="12"
        >
          <HostCard
            v-if="cpuinfo"
            color="info"
            icon="mdi-hamburger"
            :title="$t('cpudes')"
            :value="$t('cpu')"
            sub-icon="mdi-clock"
            sub-text="Just Updated"
          >
            <template v-slot:service>
               <v-card>
                <v-tabs
                  v-model="cputab"
                  background-color="grey accent-4"
                  center-active
                  dark
                >
                <v-tabs-slider color="yellow"></v-tabs-slider>
                <v-tab
                  v-for="(cpu,index) in cpuinfo"
                  :key="index"
                  :href="'#cputab-' + index"
                  class="text-capitalize"
                >
                  {{$t('cpu')+":"+index}}
                </v-tab>
                </v-tabs>
                <v-tabs-items v-model="cputab">
                  <v-tab-item
                    v-for="(cpu,index) in cpuinfo"
                    :key="index"
                    :value="'cputab-' + index"
                  >
                    <v-card>
                      <v-tabs
                        v-model="cputabindex"
                        color="cyan"
                        dark
                        flat
                      >
                       <v-tabs-slider color="yellow"></v-tabs-slider>
                        <v-tab :href="'#cputabindex-0'">
                          {{$t('base')}}
                        </v-tab>
                        <v-tab :href="'#cputabindex-1'">
                          {{$t('dashboard')}}
                        </v-tab>
                        <v-tab :href="'#cputabindex-2'">
                          {{$t('trend')}}
                        </v-tab>
                      </v-tabs>
                      <v-tabs-items v-model="cputabindex">
                        <v-tab-item :value="'cputabindex-0'">
                          <v-card height="630">
                              <v-simple-table>
                                <template v-slot:default>
                                <tbody>
                                    <tr><td>{{$t('cpuindex')}}</td><td>{{index}}</td></tr>
                                    <tr><td>{{$t('cpuvendorid')}}</td><td>{{cpu.VendorID}}</td></tr>
                                    <tr><td>{{$t('cpufamily')}}</td><td>{{cpu.Family}}</td></tr>
                                    <tr><td>{{$t('cpumodel')}}</td><td>{{cpu.Model}}</td></tr>
                                    <tr><td>{{$t('cpustepping')}}</td><td>{{cpu.Stepping}}</td></tr>
                                    <tr><td>{{$t('cpuphysicalid')}}</td><td>{{cpu.PhysicalID}}</td></tr>
                                    <tr><td>{{$t('cpucoreid')}}</td><td>{{cpu.CoreID}}</td></tr>
                                    <tr><td>{{$t('cpucores')}}</td><td>{{cpu.Cores}}</td></tr>
                                    <tr><td>{{$t('cpumodelname')}}</td><td>{{cpu.ModelName}}</td></tr>
                                    <tr><td>{{$t('cpumhz')}}</td><td>{{cpu.Mhz}}</td></tr>
                                    <tr><td>{{$t('cpucachesize')}}</td><td>{{cpu.CacheSize}}</td></tr>
                                    <tr><td>{{$t('cpuflags')}}</td><td>{{cpu.Flags}}</td></tr>     
                                    <tr><td>{{$t('cpumicrocode')}}</td><td>{{cpu.Microcode}}</td></tr>                              
                                </tbody>
                                </template>
                            </v-simple-table>
                          </v-card>
                        </v-tab-item>
                        <v-tab-item :value="'cputabindex-1'">
                            <ScoringDashboard v-if="hostmonitordata" :height="500"  Scorename="剩余Cpu" :Score="hostmonitordata.CurrCpuPer">></ScoringDashboard>
                        </v-tab-item>
                        <v-tab-item :value="'cputabindex-2'">
                            <LineChart v-if="hostmonitordata" :height="500" :keys="hostmonitordata.Keys" :values="hostmonitordata.Cpu"></LineChart>                     
                        </v-tab-item>
                      </v-tabs-items>
                    </v-card>
                  </v-tab-item>
                </v-tabs-items>
              </v-card>
            </template>
          </HostCard>
        </v-col>

        <v-col
          cols="12"
        >
          <HostCard
            v-if="memoryinfo"
            color="info"
            icon="mdi-hamburger"
            :title="$t('memorydes')"
            :value="$t('memory')"
            sub-icon="mdi-clock"
            sub-text="Just Updated"
          >
            <template v-slot:service>
                                  <v-card>
                      <v-tabs
                        v-model="memoryindex"
                        color="cyan"
                        dark
                        flat
                      >
                       <v-tabs-slider color="yellow"></v-tabs-slider>
                        <v-tab :href="'#memoryindex-0'">
                          {{$t('base')}}
                        </v-tab>
                        <v-tab :href="'#memoryindex-1'">
                          {{$t('dashboard')}}
                        </v-tab>
                        <v-tab :href="'#memoryindex-2'">
                          {{$t('trend')}}
                        </v-tab>
                      </v-tabs>
                      <v-tabs-items v-model="memoryindex">
                        <v-tab-item :value="'memoryindex-0'">
                          <v-card height="480">
                              <v-simple-table>
                                <template v-slot:default>
                                <tbody>
                                    <tr><td>{{$t('memorytotal')}}</td><td>{{memoryinfo.Total | memoryconver}}</td></tr>
                                    <tr><td>{{$t('memoryavailable')}}</td><td>{{memoryinfo.Available | memoryconver}}</td></tr>
                                    <tr><td>{{$t('memoryused')}}</td><td>{{memoryinfo.Used | memoryconver}}</td></tr>
                                    <tr><td>{{$t('memoryusedpercent')}}</td><td>{{memoryinfo.UsedPercent+"%"}}</td></tr>
                                    <tr><td>{{$t('memoryfree')}}</td><td>{{memoryinfo.Free | memoryconver}}</td></tr>
                                    <tr><td>{{$t('memoryactive')}}</td><td>{{memoryinfo.Active}}</td></tr>
                                    <tr><td>{{$t('memoryinactive')}}</td><td>{{memoryinfo.Inactive}}</td></tr>
                                    <tr><td>{{$t('memorywired')}}</td><td>{{memoryinfo.Wired}}</td></tr>
                                    <tr><td>{{$t('memorylaundry')}}</td><td>{{memoryinfo.Laundry}}</td></tr>                            
                                </tbody>
                                </template>
                            </v-simple-table>
                          </v-card>
                        </v-tab-item>
                        <v-tab-item :value="'memoryindex-1'">
                            <ScoringDashboard v-if="hostmonitordata" :height="500" Scorename="剩余内存" :Score="hostmonitordata.CurrMemoryPer"></ScoringDashboard>
                        </v-tab-item>
                        <v-tab-item :value="'memoryindex-2'">
                            <LineChart v-if="hostmonitordata" :height="500" :keys="hostmonitordata.Keys" :values="hostmonitordata.Memory"></LineChart>                     
                        </v-tab-item>
                      </v-tabs-items>
                  </v-card>
            </template>
          </HostCard>
        </v-col>
    </div>
</template>

<script>
import HostCard from "./components/MaterialHostCard"
import {ScoringDashboard,LineChart} from "@/components/echarts"
// import PieChart from "@/components/PieChart"
// import LineChart from "@/components/LineChart"
import moment from 'moment'
import { mapGetters } from 'vuex'
export default {
    name:"Base",
    components:{
        HostCard,
        ScoringDashboard,
        LineChart,
        // PieChart,
        // LineChart
    },
    computed: {
      ...mapGetters([
        'hostinfo',
        'cpuinfo',
        'memoryinfo',
        'hostmonitordata',
      ]),

    },
    data :function(){
      return {
        cputab:null,
        cputabindex:null,
        memoryindex:null,
      }
    },
    filters: {
        formatDate(time) {
            var date = new Date(time*1000);
            return moment(date).format("YYYY-MM-DD HH:mm:ss");
        },
        //秒转化成 时分秒
        secondToDate(result) {
            var h = Math.floor(result / 3600);
            var m = Math.floor((result / 60 % 60));
            var s = Math.floor((result % 60));
            return result = h + ":" + m + ":" + s;
        },
        memoryconver(limit){
          var size = "";
          if( limit < 0.1 * 1024 ){ //如果小于0.1KB转化成B
            size = limit.toFixed(2) + "B"; 	
          }else if(limit < 0.1 * 1024 * 1024 ){//如果小于0.1MB转化成KB
            size = (limit / 1024).toFixed(2) + "KB";			
          }else if(limit < 0.1 * 1024 * 1024 * 1024){ //如果小于0.1GB转化成MB
            size = (limit / (1024 * 1024)).toFixed(2) + "MB";
          }else{ //其他转化成GB
            size = (limit / (1024 * 1024 * 1024)).toFixed(2) + "GB";
          }
          
          var sizestr = size + ""; 
          var len = sizestr.indexOf(".");
          var dec = sizestr.substr(len + 1, 2);
          if(dec == "00"){//当小数点后为00时 去掉小数部分
            return sizestr.substring(0,len) + sizestr.substr(len + 3,2);
          }
          return sizestr;
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
      this.$store.dispatch('console/gethostmonitordata')
    },
    beforeDestroy () {
      //  this.stoptimer()
    }
}
</script>

<style lang="sass">

.legoinfo
  width: 100%
  height: 100%
  background-color: #EFEBE9
  border-radius: 5px
  .infotitle
    padding-top: 10px 
    padding-left: 10px
    padding-bottom: 0
    font-size: 20px
    height: 20px
    color: #3E2723
  .infotitle-sub
    padding-left: 10px
    padding-bottom: 0
    font-size: 10px
    height: 10px

  .legoinfo-centext
    padding: 10px
    .legotag
      color: #6D4C41
    .legodes
      min-height: 100px
      font-size: 10px
      color: #A1887F
  .legoinfo-bottom
    font-size: 10px
    color: #A1887F
    display: flex
    justify-content: space-between


.serviceinfo
  width: 100%
  height: 100%
  background-color: #ECEFF1
  border-radius: 5px
  .infotitle
    padding-top: 10px 
    padding-left: 10px
    padding-bottom: 0
    font-size: 20px
    height: 20px
    color: #263238
  .infotitle-sub
    padding-left: 10px
    padding-bottom: 0
    font-size: 10px
    height: 10px
    color: #90A4AE
  .infotitle-name
    font-size: 15px
    color: #78909C
    text-align: center
  .infotitle-value
    font-size: 15px
    color: #B0BEC5

.cpuinfo
  width: 100%
  height: 100%
  background-color: #FFF3E0
  border-radius: 5px
  .infotitle
    font-weight: bold
    color: #E65100
  .infotitle-sub
    padding-left: 10px
    padding-bottom: 0
    font-size: 10px
    height: 10px
    color: #90A4AE
  .legoinfo-centext
    padding: 10px
    background-color: #FFF3E0
    .infotitle-name
      font-size: 15px
      color: #F57C00
      text-align: center
    .infotitle-value
      font-size: 15px
      color: #FFCC80


.memoryinfo
  width: 100%
  height: 100%
  background-color: #FAFAFA
  border-radius: 5px
  .infotitle
    padding-top: 10px 
    padding-left: 10px
    padding-bottom: 0
    font-size: 20px
    height: 20px
    color: #263238
  .infotitle-sub
    padding-left: 10px
    padding-bottom: 0
    font-size: 10px
    height: 10px
    color: #90A4AE
  .legoinfo-centext
    .infotitle-name
      font-size: 15px
      color: #F57C00
      text-align: center
    .infotitle-value
      font-size: 15px
      color: #FFCC80

</style>
