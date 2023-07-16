<template>
  <main>
    <el-form  :model="form" class="demo-form-inline">
      <el-form-item label="键">
        <el-input v-model="form.key" disabled placeholder=""  />
      </el-form-item>
      <el-row>
        <el-col :span="11">
          <el-form-item label="过期时间">
            <el-input v-model="form.ttl"  placeholder=""  />
          </el-form-item>
        </el-col>
        <el-col :span="2">
        </el-col>
        <el-col span="11">
          <el-form-item label="数据类型">
            <el-input v-model="form.type" disabled placeholder=""  />
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="值">
        <el-input v-model="form.value"  placeholder=""  />
      </el-form-item>
      <el-form-item >
        <el-button type="primary" @click="updateKey"> 保存</el-button>
      </el-form-item>
<!--      <el-form-item>-->
<!--        <el-button type="primary" @click="onSubmit">查询</el-button>-->
<!--      </el-form-item>-->
    </el-form>
  </main>
</template>

<script setup>
import {ref, watch} from 'vue'
import {GetKeyValue, UpdataKeyValue} from "../../wailsjs/go/main/App.js";
import {ElNotification} from "element-plus";
let props=defineProps(['keyDb','keyConnIdentity','KeyKey'])
let form = ref({})
watch(()=>props.KeyKey,()=>{
  getKeyValue()
})

function  getKeyValue(){
  GetKeyValue({conn_identity:props.keyConnIdentity,db:props.keyDb,key:props.KeyKey}).then(res=>{
    if (res.code!==200){
      ElNotification({
        title:res.msg,
        type:"error",
      })
      return
    }
    form.value=res.data
    form.value.key=props.KeyKey
  })
}

function updateKey() {
  UpdataKeyValue({conn_identity:props.keyConnIdentity,db:props.keyDb,key:props.KeyKey,value:form.value.value,ttl:form.value.ttl}).then(res=>{
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
    getKeyList()
  })
}

</script>

<style scoped>

</style>