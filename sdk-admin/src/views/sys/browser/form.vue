<script setup lang="ts">
import { ref, reactive } from "vue";
import type { FormRules } from "element-plus";
import { BrowserFormProps } from "@/api/sys/browser";

const props = withDefaults(defineProps<BrowserFormProps>(), {
  formInline: () => ({
    id: 0,
    name: "",
    envId: 0,
    userId: 0,
    data: ""
  })
});

/** 自定义表单规则校验 */
const formRules = reactive(<FormRules>{
  name: [
    { required: true, message: "名称为必填项", trigger: "blur" },
    { min: 1, max: 255, message: "长度在 1 到 255 个字符", trigger: "blur" }
  ],
  envId: [{ required: true, message: "环境ID为必填项", trigger: "blur" }],
  userId: [{ required: true, message: "用户ID为必填项", trigger: "blur" }],
  data: [{ required: true, message: "数据为必填项", trigger: "blur" }]
});

const ruleFormRef = ref();
const newFormInline = ref({ ...props.formInline });

function getRef() {
  return ruleFormRef.value;
}

defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="100px"
  >
    <el-form-item label="主键" prop="id">
      <el-input
        v-model.number="newFormInline.id"
        clearable
        disabled
        placeholder="请输入主键"
      />
    </el-form-item>
    <el-form-item label="名称" prop="name">
      <el-input
        v-model="newFormInline.name"
        clearable
        placeholder="请输入名称"
      />
    </el-form-item>
    
    <el-form-item label="环境ID" prop="envId">
      <el-input-number
        v-model="newFormInline.envId"
        placeholder="请输入环境ID"
        :min="0"
        controls-position="right"
        style="width: 100%"
      />
    </el-form-item>
    
    <el-form-item label="用户ID" prop="userId">
      <el-input-number
        v-model="newFormInline.userId"
        placeholder="请输入用户ID"
        :min="0"
        controls-position="right"
        style="width: 100%"
      />
    </el-form-item>
    
    <el-form-item label="数据" prop="data">
      <el-input
        v-model="newFormInline.data"
        type="textarea"
        :rows="4"
        placeholder="请输入数据"
        clearable
      />
    </el-form-item>
  </el-form>
</template>