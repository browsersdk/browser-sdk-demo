<script setup lang="ts">
import { ref, onMounted } from "vue";
import ReCol from "@/components/ReCol";
import { formRules } from "../utils/rule";
import { AdminFormProps } from "@/api/sys/admin";
import { usePublicHooks } from "@/utils/hooks";
import { useMember } from "../utils/hook";
import { getAdminRolePage } from "@/api/sys/admin-role";
import { getAdmin } from "@/api/sys/admin";
const props = withDefaults(defineProps<AdminFormProps>(), {
  formInline: () => ({
    type: "",
    higherDeptOptions: [],
    id: 0,
    nickname: null,
    name: null,
    phone: null,
    deptPath: null,
    deptId: 0,
    postId: 0,
    status: null,
    gender: null,
    birthday: null,
    username: null,
    email: null,
    password: null,
    avatar: null,
    bio: null,
    roleId: 0,
    remark: null,
    lockTime: null,
    clientId: null,
    regIp: null,
    ipLocation: null
  })
});

//const searchRef = ref();
const ruleFormRef = ref();
const { switchStyle } = usePublicHooks();
const newFormInline = ref(props.formInline);

const treeRef = ref();
const tableRef = ref();

const roleList = ref([]);

const { genderOptions } = useMember(tableRef, treeRef);

function getRef() {
  return ruleFormRef.value;
}
const getRolePage = async () => {
  const { data } = await getAdminRolePage({
    page: 1,
    pageSize: 10000
  });
  roleList.value = data.list || [];
  // roleList.value.push({
  //   id: -1,
  //   name: "超管"
  // });
  // if (roleList.value && roleList.value) {
  //   roleList.value.forEach(item => {
  //     item.label = item.name;
  //     item.id = item.id;
  //   });
  // }
};
getRolePage();
const getDetail = () => {
  getAdmin({ id: newFormInline.value.id }).then(res => {
    if (res.code === 200) {
      newFormInline.value.id = res.data.id;
      newFormInline.value.name = res.data.name;
      newFormInline.value.nickname = res.data.nickname;
      newFormInline.value.phone = res.data.phone;
      newFormInline.value.email = res.data.email;
      newFormInline.value.username = res.data.username;
      newFormInline.value.birthday = res.data.birthday;
      newFormInline.value.gender = res.data.gender;
      newFormInline.value.deptId = res.data.deptId;
      newFormInline.value.roleId = res.data.roleId;
      newFormInline.value.status = res.data.status;
    }
  });
};
onMounted(() => {
  getDetail();
});
defineExpose({ getRef });
</script>

<template>
  <el-form
    ref="ruleFormRef"
    :model="newFormInline"
    :rules="formRules"
    label-width="82px"
  >
    <el-row :gutter="30">
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="姓名" prop="name">
          <el-input
            v-model="newFormInline.name"
            clearable
            placeholder="请输入用户名称"
          />
        </el-form-item>
      </re-col>
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="昵称" prop="nickname">
          <el-input
            v-model="newFormInline.nickname"
            clearable
            placeholder="请输入用户昵称"
          />
        </el-form-item>
      </re-col>
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="手机号" prop="phone">
          <el-input
            v-model="newFormInline.phone"
            clearable
            placeholder="请输入手机号"
          />
        </el-form-item>
      </re-col>
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="邮箱" prop="email">
          <el-input
            v-model="newFormInline.email"
            clearable
            placeholder="请输入邮箱"
          />
        </el-form-item>
      </re-col>
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="用户名" prop="username">
          <el-input
            v-model="newFormInline.username"
            clearable
            placeholder="请输入用户名"
          />
        </el-form-item>
      </re-col>
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="密码" prop="password">
          <el-input
            v-model="newFormInline.password"
            clearable
            placeholder="请输入密码"
            type="password"
          />
        </el-form-item>
      </re-col>

      <!-- <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="入职时间" prop="entryTime">
          <el-date-picker
            style="width: 100%"
            v-model="newFormInline.entryTime"
            type="date"
            placeholder="选择日期"
          />
        </el-form-item>
      </re-col>

      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="离职时间" prop="retireTime">
          <el-date-picker
            style="width: 100%"
            v-model="newFormInline.retireTime"
            type="date"
            placeholder="选择日期"
          />
        </el-form-item>
      </re-col> -->
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="生日" prop="birthday">
          <el-date-picker
            style="width: 100%"
            v-model="newFormInline.birthday"
            type="date"
            placeholder="选择日期"
          />
        </el-form-item>
      </re-col>

      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="性别">
          <el-select
            v-model="newFormInline.gender"
            placeholder="请选择用户性别"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in genderOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </re-col>

      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="归属部门">
          <el-cascader
            class="w-full"
            v-model="newFormInline.deptId"
            :options="newFormInline.higherDeptOptions"
            :props="{
              value: 'id',
              label: 'name',
              emitPath: false,
              checkStrictly: true
            }"
            clearable
            filterable
            placeholder="请选择归属部门"
          >
            <template #default="{ node, data }">
              <span>{{ data.name }}</span>
              <span v-if="!node.isLeaf"> ({{ data.children.length }}) </span>
            </template>
          </el-cascader>
        </el-form-item>
      </re-col>
      <!-- <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="职位标签" prop="postTag">
          <el-select
            v-model="newFormInline.postId"
            placeholder="请选择职位"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in postOptions"
              :key="index"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>
      </re-col> -->
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="角色" prop="roleId">
          <el-select
            v-model="newFormInline.roleId"
            placeholder="请选择角色"
            class="w-full"
            clearable
          >
            <el-option
              v-for="(item, index) in roleList"
              :key="index"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
      </re-col>
      <re-col :value="8" :xs="24" :sm="24">
        <el-form-item label="用户状态">
          <el-switch
            v-model="newFormInline.status"
            inline-prompt
            :active-value="1"
            :inactive-value="2"
            active-text="在职"
            inactive-text="离职"
            :style="switchStyle"
          />
        </el-form-item>
      </re-col>
    </el-row>
  </el-form>
</template>
