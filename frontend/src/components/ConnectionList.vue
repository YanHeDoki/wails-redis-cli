<template>

  <main>
    <div class="demo-collapse">
      <el-collapse accordion>
        <el-collapse-item v-for="item in list" :name="item.identity" @click=getInfo(item.identity)>
         <template  #title>
           <div class="item">
             <div>
                    {{ item.name }}
             </div>
             <div style="display: flex">
               <ConnectionManage @click.stop  title="编辑" btn-type="text" :data="item" />
                <el-popconfirm title="确认删除吗?" @confirm="connectionDelete(item.identity)">
                  <template #reference>
                    <el-button link type="danger" @click.stop>删除</el-button>
                  </template>
                </el-popconfirm>
             </div>
           </div>
         </template>
          <div v-for="db in infoDbList" @click="selectDB(db.key,item.identity)">
            <div v-if="db.key!==selectDBKey" class="db_item">
              {{db.key}}( {{db.number}})
            </div>
            <div v-else class="select_item">
              {{db.key}}( {{db.number}})
            </div>
          </div>
        </el-collapse-item>
      </el-collapse>
    </div>
  </main>
</template>

<script setup>
import {ref,watch} from "vue";
import {ConnectionDelete, ConnectionList, DbList} from "../../wailsjs/go/main/App.js";
import ConnectionManage from "./ConnectionManage.vue";
import {ElNotification} from "element-plus";

let list=ref()
let props= defineProps(['flush'])
let infoDbList=ref()
let selectDBKey=ref()
let emits=defineEmits(['emit-select-db'])
watch(props,(newFlush)=>{
  connectionList()
})
function connectionList() {
  ConnectionList().then(res=>{
    if (res.code!=200){
      ElNotification({
        title:res.msg,
        type:res.error,
      })
      return
    }
    list.value=res.data
  })
}
connectionList()

function connectionDelete(identity) {
  ConnectionDelete(identity).then(res=>{
    if (res.code!==200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
      return
    }
    ElNotification({
      title:res.msg,
      type:"success",
    })
    connectionList()
  })
}
function getInfo(identity){
  //获取数据库列表
  DbList(identity).then(res=>{
    if (res.code!==200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
      return
    }
    // ElNotification({
    //   title:res.msg,
    //   type:"success",
    // })
    infoDbList.value=res.data
  })
}
function selectDB(db,connIdentity){
  selectDBKey.value=db
  var number = Number(db.substring(2));
  emits('emit-select-db',number,connIdentity)
}

</script>
<style scoped>


</style>