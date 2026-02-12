import dayjs from "dayjs";
import editForm from "../form.vue";
import { message } from "@/utils/message";
import {
  type BrowserFormItemProps,
  getBrowserPage,
  createBrowser,
  updateBrowser,
  delBrowser
} from "@/api/sys/browser";
import { addDialog } from "@/components/ReDialog";
import { type PaginationProps } from "@pureadmin/table";
import { reactive, ref, onMounted, h, toRaw } from "vue";

export function useBrowser() {
  const qform = reactive({
    page: 1,
    pageSize: 10,
    name: null,
    envId: null,
    userId: null
  });
  const formRef = ref();
  const dataList = ref([]);
  const loading = ref(true);
  const pagination = reactive<PaginationProps>({
    total: 0,
    pageSize: 10,
    currentPage: 1,
    background: true
  });
  const columns: TableColumnList = [
    {
      label: "主键",
      prop: "id",
      minWidth: 120
    },
    {
      label: "名称",
      prop: "name",
      minWidth: 120
    },
    {
      label: "环境ID",
      prop: "envId",
      minWidth: 120
    },
    {
      label: "用户ID",
      prop: "userId",
      minWidth: 120
    },
    {
      label: "数据",
      prop: "data",
      minWidth: 200,
      showOverflowTooltip: true
    },
    {
      label: "创建时间",
      prop: "createdAt",
      minWidth: 180,
      formatter: ({ createdAt }) =>
        dayjs(createdAt).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "最后更新时间",
      prop: "updatedAt",
      minWidth: 180,
      formatter: ({ updatedAt }) =>
        dayjs(updatedAt).format("YYYY-MM-DD HH:mm:ss")
    },
    {
      label: "操作",
      width: 140,
      slot: "operation"
    }
  ];

  function handleDelete(row) {
    delBrowser({ ids: [row.id] }).then(res => {
      if (res.code == 200) {
        message(`删除成功`, { type: "success" });
        onSearch();
      } else {
        message(`删除失败`, { type: "error" });
      }
    });
  }

  function handleSizeChange(val: number) {
    qform.pageSize = val;
    onSearch();
  }

  function handleCurrentChange(val: number) {
    qform.page = val;
    onSearch();
  }

  function handleSelectionChange(val) {
    console.log("handleSelectionChange", val);
  }

  async function onSearch() {
    loading.value = true;
    const { data } = await getBrowserPage(toRaw(qform));
    dataList.value = data.list;
    pagination.total = data.total;
    pagination.pageSize = data.pageSize;
    pagination.currentPage = data.currentPage;

    setTimeout(() => {
      loading.value = false;
    }, 500);
  }

  const resetForm = formEl => {
    if (!formEl) return;
    formEl.resetFields();
    onSearch();
  };

  function openDialog(title = "新增", row?: BrowserFormItemProps) {
    addDialog({
      title: `${title}浏览器`,
      props: {
        formInline: {
          id: row?.id ?? 0,
          name: row?.name ?? "",
          envId: row?.envId ?? 0,
          userId: row?.userId ?? 0,
          data: row?.data ?? ""
        }
      },
      width: "48%",
      draggable: true,
      fullscreenIcon: true,
      closeOnClickModal: false,
      contentRenderer: () => h(editForm, { 
        ref: formRef, 
        formInline: {
          id: 0,
          name: "",
          envId: 0,
          userId: 0,
          data: ""
        } 
      }),
      beforeSure: (done, { options }) => {
        const FormRef = formRef.value.getRef();
        const curData = options.props.formInline as BrowserFormItemProps;
        FormRef.validate(valid => {
          if (valid) {
            // 表单规则校验通过
            if (title === "新增") {
              createBrowser(curData).then(res => {
                if (res.code == 200) {
                  message(res.msg, {
                    type: "success"
                  });
                  onSearch(); // 刷新表格数据
                } else {
                  message(res.msg, {
                    type: "error"
                  });
                }
              });
            } else {
              updateBrowser(curData).then(res => {
                if (res.code == 200) {
                  message(res.msg, {
                    type: "success"
                  });
                  onSearch(); // 刷新表格数据
                } else {
                  message(res.msg, {
                    type: "error"
                  });
                }
              });
            }
            done(); // 关闭弹框
          }
        });
      }
    });
  }

  onMounted(() => {
    onSearch();
  });

  return {
    qform,
    loading,
    columns,
    dataList,
    pagination,
    onSearch,
    resetForm,
    openDialog,
    handleDelete,
    handleSizeChange,
    handleCurrentChange,
    handleSelectionChange
  };
}