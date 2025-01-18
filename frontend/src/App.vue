<script lang="ts" setup>
import { defineComponent, onMounted, onUnmounted, ref, watch } from 'vue';
import { GetSerialNames, PortClose, PortOpen } from '../wailsjs/go/main/App';
import HelloWorld from './components/HelloWorld.vue' 


const listName = ref<string[]>();
const baudRate = ref<string>();
const selectedName = ref<string>();

const textTerminal = ref<string>("");

function sleep(n:number){
  return new Promise((r,x)=>{
    setTimeout(r, n);
  })
}


async function getSerialList() {
  listName.value = await GetSerialNames();
}
 

onMounted(() => {
  PortClose();
  getSerialList(); 

  (window as any).runtime.EventsOn("printToweb", (data:number[])=>{
    var textDecoder = new TextDecoder("utf-8");
    var str = textDecoder.decode(new Uint8Array(data));
    textTerminal.value += str;
  });


  (async ()=>{
    while(true){
      await sleep(1000);
      getSerialList();
    }

  })();
})

function startConnect(){
  if(!baudRate.value || baudRate.value == ""){
     baudRate.value = "9600"
  }

  var n = Number(baudRate.value);
  if(isNaN(n)) return;

  if(selectedName.value){

    PortOpen(selectedName.value, n, 100);
  }

  
}

onUnmounted(()=>{
  PortClose();

})

</script>

<template>

  <div class="mainpage">
    <h3>Serial Reader</h3>

    <table>
      <tbody> 
        <tr>
          <td class="note">Serial Name</td>
          <td> : </td>
          <td>

            <select v-model="selectedName">
              <option v-for="item in listName">{{ item }}</option>
            </select>
          </td>
        </tr>

        <tr>
          <td class="note">baudRate</td>
          <td> : </td>
          <td>

            <input v-model="baudRate"  type="number" />
          </td>
        </tr>
        <tr>
          <td colspan="3">
            
            <button :onclick="startConnect">Start Connect</button>
          </td>
        </tr>
      </tbody>
    </table>

    <textarea :value="textTerminal"></textarea>
  </div>
</template>

<style scoped>
.mainpage {
  text-align: center;
}

.mainpage table {
  background-color: #283850;
  ;
  margin: auto;
  width: 100%;
  border-collapse: collapse;
}

.mainpage table .note {
  text-align: left;
}

.mainpage table td {
  border: solid 1px #ccc;
  padding: 3px;
}
.mainpage table td input{
  box-sizing: border-box;
  width: 100%;
  min-height: 30px;
}

.mainpage table td button{
  margin: 10px;
  background-color: #32CD32;
  border: none;
  min-height: 40px;
  min-width: 100px;
  border-radius: 5px;
  cursor: pointer;
}
.mainpage table td button:active { 
  transform: scale(0.97);
}


.mainpage select{
  width: 100%;
  min-height: 30px;
}
.mainpage textarea{
  width: 100%;
  height: 200px;
  margin-top: 20px;
  background-color: black;
  color: #ccc;
}
</style>
