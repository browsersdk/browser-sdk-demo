<script setup lang="ts">
import { ref, reactive } from "vue";
import type { FormRules } from "element-plus";
import { CounterFormProps } from "@/api/sys/counter";

const props = withDefaults(defineProps<CounterFormProps>(), {
  formInline: () => ({
    key: null,
    val: 0
  })
});

/** 自定义表单规则校验 */
const formRules = reactive(<FormRules>{
  name: [{ required: true, message: "名称为必填项", trigger: "blur" }]
});

const ruleFormRef = ref();
const newFormInline = ref(props.formInline);

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
    label-width="82px"
  >
    <el-form-item label="关键字" prop="key">
      <el-input
        v-model="newFormInline.key"
        clearable
        placeholder="请输入关键字"
      />
    </el-form-item>
    <el-form-item label="值" prop="seq">
      <el-input
        v-model.number="newFormInline.seq"
        clearable
        placeholder="请输入值"
      />
    </el-form-item>
  </el-form>
</template>
