<script lang="ts" setup>
import { defineComponent, onMounted, onUnmounted, ref, watch } from 'vue';
import { GetSerialNames, IsConnected, PortClose, PortOpen } from '../wailsjs/go/main/App';
import HelloWorld from './components/HelloWorld.vue' 

export type PropType = {
  openAbout :()=>void
}
const prop = defineProps<PropType>();

const listName = ref<string[]>();
const baudRate = ref<string>();
const selectedName = ref<string>();
const isPortsConnected = ref<boolean>(false);

const textTerminal = ref<string>("");

function sleep(n:number){
  return new Promise((r,x)=>{
    setTimeout(r, n);
  })
}


async function getSerialList() {
  listName.value = await GetSerialNames();
}

var isStillOpen = false;

onMounted(() => {
  

  isStillOpen = true;
  PortClose();
  getSerialList(); 

  (window as any).runtime.EventsOn("printToweb", (data:number[])=>{
    var textDecoder = new TextDecoder("utf-8");
    var str = textDecoder.decode(new Uint8Array(data));
    textTerminal.value += str;
  });


  (async ()=>{
    while(isStillOpen){
      await sleep(1000);
      getSerialList();
      isPortsConnected.value = await IsConnected();
      console.log("looping " + isPortsConnected.value );
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
  isStillOpen = false;
})

</script>

<template>

  <div class="mainpage">
    <div>

      <h3>Serial Reader</h3>
    </div>

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
          <td class="note">BaudRate</td>
          <td> : </td>
          <td>

            <input v-model="baudRate"  type="number" />
          </td>
        </tr>
        
        <tr>
          <td></td>
          <td></td>
          <td>
            <button v-if="!isPortsConnected" :onclick="startConnect">Start Connect</button> 
            <button class="btnred" v-if="isPortsConnected" :onclick="()=>{ PortClose(); }">Disconnect</button> 
          </td>
        </tr>
         
      </tbody>
    </table>

    <textarea :value="textTerminal"></textarea>
     
  </div>
  <button class="btnontop" v-on:click="prop.openAbout()"><i class="btnAbout"></i></button>
</template>

<style scoped>
.mainpage {
  text-align: center;
  display: flex;
  flex-direction: column; /* Atau column, sesuai kebutuhan */
  justify-content: space-between; /* Atur penyebaran elemen */
  align-items: center; /* Vertikal alignment */
  height: 100%;  
  gap: 5px;
}
.mainpage h3{
  margin: 0px;
}
.mainpage table { 
  margin: auto;
  width: 100%;
  border-collapse: collapse;
}

.mainpage table .note {
  text-align: left;
}

.mainpage table td { 
  padding: 3px;
}
.mainpage table td input{
  box-sizing: border-box;
  width: 100%;
  min-height: 30px;
}

.mainpage table td  button{ 
  background-color: #0525de;
  border: none; 
  min-width: 100px;
  width: 100%;
  border-radius: 5px;
  cursor: pointer;
  color: white;
  padding: 5px;
}
.mainpage table td  button:active { 
  transform: scale(0.97);
}
.mainpage button.btnred{
  background-color: red;
}


.mainpage table td  select{
  width: 100%;
  min-height: 30px;
}
.mainpage textarea{
  box-sizing: border-box;
  width: 100%; 
  height: 100%;
  margin-top: 20px;
  background-color: black;
  color: #ccc;
  resize: none;
  flex: 1;
}

.btnAbout{
  box-sizing: border-box;
  display: block;
  width: 100%;
  height: 100%;
  background-image: url('data:image/svg+xml,%0A%20%20%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2216%22%20height%3D%2216%22%20fill%3D%22currentColor%22%20class%3D%22bi%20bi-blockquote-left%22%20viewBox%3D%220%200%2016%2016%22%3E%0A%20%20%3Cpath%20d%3D%22M2.5%203a.5.5%200%200%200%200%201h11a.5.5%200%200%200%200-1zm5%203a.5.5%200%200%200%200%201h6a.5.5%200%200%200%200-1zm0%203a.5.5%200%200%200%200%201h6a.5.5%200%200%200%200-1zm-5%203a.5.5%200%200%200%200%201h11a.5.5%200%200%200%200-1zm.79-5.373q.168-.117.444-.275L3.524%206q-.183.111-.452.287-.27.176-.51.428a2.4%202.4%200%200%200-.398.562Q2%207.587%202%207.969q0%20.54.217.873.217.328.72.328.322%200%20.504-.211a.7.7%200%200%200%20.188-.463q0-.345-.211-.521-.205-.182-.568-.182h-.282q.036-.305.123-.498a1.4%201.4%200%200%201%20.252-.37%202%202%200%200%201%20.346-.298zm2.167%200q.17-.117.445-.275L5.692%206q-.183.111-.452.287-.27.176-.51.428a2.4%202.4%200%200%200-.398.562q-.165.31-.164.692%200%20.54.217.873.217.328.72.328.322%200%20.504-.211a.7.7%200%200%200%20.188-.463q0-.345-.211-.521-.205-.182-.568-.182h-.282a1.8%201.8%200%200%201%20.118-.492q.087-.194.257-.375a2%202%200%200%201%20.346-.3z%22%2F%3E%0A%20%20%3C%2Fsvg%3E%20%20%0A%20%20');
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
  
}

.btnontop{
  display: block;
  position: absolute;
  top: 10px;
  left: 10px;
  width: 30px;
  height: 30px;
  background-color: #ffffff;
  border: none;
  border-radius: 10px;
  cursor: pointer; 
}
.btnontop:active{
  transform: scale(0.80); 
}
</style>
