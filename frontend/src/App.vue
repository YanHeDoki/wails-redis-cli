<script setup xmlns:el-col="http://www.w3.org/1999/html">

import ConnectionList from "./components/ConnectionList.vue";
import ConnectionManage from "./components/ConnectionManage.vue";
import {ConnectionCreate, CreateKeyValue, DbList, GetKeyValue, KeyList} from "../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
import {ref} from "vue";
import Keys from "./components/keys.vue";
import KeyValue from "./components/KeyValue.vue";

let flushFlag=ref(true)
let keyValue=ref()

let keyDb=ref()
let keyConnIdentity=ref()

let KeyKey=ref()

function flushConnectionList() {
  flushFlag.value=!flushFlag.value
}
// GetKeyValue({conn_Identity: "e04dc1a7-f38c-45b7-89cb-51806f7a5d14", db: 0,key:"k2"}).then(res=>{
//   keyValue.value=res
// })
function selectDB(db,connIdentity) {
  keyDb.value=db
  keyConnIdentity.value=connIdentity
}

function selectKey(key) {
  KeyKey.value=key
}
function addKeyValue() {
  CreateKeyValue({conn_identity:"e04dc1a7-f38c-45b7-89cb-51806f7a5d14",db:0,key:"k2",type:"string"})
}


</script>

<template>
 <el-row>
   <el-col :span="5" style="height: 100vh;padding: 12px">
     <div  style="margin-bottom: 12px">
       <ConnectionManage title="新建连接" btn-type="primary" @emit-connection-List="flushConnectionList"></ConnectionManage>
       <el-button  @click="addKeyValue()">
         测试
       </el-button>
     </div>
     <ConnectionList @emit-select-db="selectDB" :flush="flushFlag"/>
   </el-col>
   <el-col :span="7" style="padding: 12px">
    <keys :keyDb="keyDb" :keyConnIdentity="keyConnIdentity" @emit-select-key="selectKey"/>
   </el-col>
  <el-col :span="12" style="padding: 12px">
   <KeyValue :keyDb="keyDb" :keyConnIdentity="keyConnIdentity" :KeyKey="KeyKey"/>
  </el-col>
 </el-row>
</template>

<style>
#logo {
  /*display: block;*/ 
  /*width: 50%;*/
  /*height: 50%;*/
  /*margin: auto;*/
  /*padding: 10% 0 0;*/
  /*background-position: center;*/
  /*background-repeat: no-repeat;*/
  /*background-size: 100% 100%;*/
  /*background-origin: content-box;*/
}
</style>
