// 这里存放本地图标，在 src/layout/index.vue 文件中加载，避免在首启动加载
import { getSvgInfo } from "@pureadmin/utils";
import { addIcon } from "@iconify/vue/dist/offline";

// https://icon-sets.iconify.design/ep/?keyword=ep
import EpHomeFilled from "~icons/ep/home-filled?raw";

// https://icon-sets.iconify.design/ri/?keyword=ri
import RiSearchLine from "~icons/ri/search-line?raw";
import RiInformationLine from "~icons/ri/information-line?raw";
import Bill from "~icons/ri/bill-line?raw";
import Team from "~icons/ri/team-line?raw";

import Lollipop from "~icons/ep/lollipop?raw";
import UbuntuFill from "~icons/ri/ubuntu-fill?raw";
import Menu from "~icons/ep/menu?raw";
import Edit from "~icons/ep/edit?raw";
import SetUp from "~icons/ep/set-up?raw";
import TerminalWindowLine from "~icons/ri/terminal-window-line?raw";
import Guide from "~icons/ep/guide?raw";
import Card from "~icons/ri/bank-card-line?raw";
import ListCheck from "~icons/ri/list-check?raw";
import Histogram from "~icons/ep/histogram?raw";
import Ppt from "~icons/ri/file-ppt-2-line?raw";
import CheckboxCircleLine from "~icons/ri/checkbox-circle-line?raw";
import FlUser from "~icons/ri/admin-line?raw";
import Role from "~icons/ri/admin-fill?raw";
import Setting from "~icons/ri/settings-3-line?raw";
import Dept from "~icons/ri/git-branch-line?raw";
import Monitor from "~icons/ep/monitor?raw";
import Customer from "~icons/ri/user-star-line?raw";
import Planet from "~icons/ri/planet-line?raw";
import Task from "~icons/ri/task-line?raw";
import CalendarCheck from "~icons/ri/calendar-check-line?raw";
import OpenArm from "~icons/ri/open-arm-line?raw";
import Table from "~icons/ri/table-line?raw";
import Generate from "~icons/ri/ai-generate?raw";
import Operation from "~icons/ep/operation?raw";
import SwapLine from "~icons/ri/swap-line?raw";
import Bell from "~icons/ep/bell?raw";
import BellFill from "~icons/ep/bell-filled?raw";
import Ai from "~icons/ri/openai-line?raw";
import Apple from "~icons/ri/apple-line?raw";
import App from "~icons/ri/apps-2-line?raw";
import Film from "~icons/ri/film-line?raw";
import Video from "~icons/ri/video-line?raw";
import Voice from "~icons/ri/voice-recognition-fill?raw";
import UserSetting from "~icons/ri/user-settings-line?raw";
import Log from "~icons/ic/outline-receipt?raw";
import Device from "~icons/ri/device-line?raw";
import Key from "~icons/ri/key-2-line?raw";
import Function from "~icons/ri/function-line?raw";
import Company from "~icons/ri/building-line?raw";
import Group from "~icons/ri/group-line?raw";
import Cmd from "~icons/ri/command-line?raw";
import Aim from "~icons/ep/aim?raw";
import Chat from "~icons/ep/chat-dot-round?raw";
import Tag from "~icons/ep/collection-tag?raw";
import Tickets from "~icons/ep/tickets?raw";
import Bot from "~icons/ri/robot-line?raw";
import No from "~icons/ri/notion-line?raw";
import Counter from "~icons/ri/upload-line?raw";
import ChatHistory from "~icons/ri/chat-history-line?raw";
import Module from "~icons/ri/sound-module-line?raw";
import ShieldUser from "~icons/ri/shield-user-line?raw";
import Drawer from "~icons/ri/archive-drawer-line?raw";
// https://icon-sets.iconify.design/ep/?keyword=ep
import EpMenu from "~icons/ep/menu?raw";
import EpEdit from "~icons/ep/edit?raw";
import EpGuide from "~icons/ep/guide?raw";
import EpSetUp from "~icons/ep/set-up?raw";
import EpMonitor from "~icons/ep/monitor?raw";
import EpLollipop from "~icons/ep/lollipop?raw";
import EpHistogram from "~icons/ep/histogram?raw";

// https://icon-sets.iconify.design/ri/?keyword=ri
import RiMindMap from "~icons/ri/mind-map?raw";
import RiAdminFill from "~icons/ri/admin-fill?raw";
import RiTableLine from "~icons/ri/table-line?raw";
import RiLinksFill from "~icons/ri/links-fill?raw";
import RiAdminLine from "~icons/ri/admin-line?raw";
import RiListCheck from "~icons/ri/list-check?raw";
import RiWindowLine from "~icons/ri/window-line?raw";
import RiUbuntuFill from "~icons/ri/ubuntu-fill?raw";
import RiHistoryFill from "~icons/ri/history-fill?raw";
import RiEditBoxLine from "~icons/ri/edit-box-line?raw";
import RiCodeBoxLine from "~icons/ri/code-box-line?raw";
import RiArtboardLine from "~icons/ri/artboard-line?raw";
import RiMarkdownLine from "~icons/ri/markdown-line?raw";
import RiFileInfoLine from "~icons/ri/file-info-line?raw";
import RiBankCardLine from "~icons/ri/bank-card-line?raw";
import RiFilePpt2Line from "~icons/ri/file-ppt-2-line?raw";
import RiGitBranchLine from "~icons/ri/git-branch-line?raw";
import RiSettings3Line from "~icons/ri/settings-3-line?raw";
import RiUserVoiceLine from "~icons/ri/user-voice-line?raw";
import RiBookmark2Line from "~icons/ri/bookmark-2-line?raw";
import RiFileSearchLine from "~icons/ri/file-search-line?raw";
import RiChatSearchLine from "~icons/ri/chat-search-line?raw";
import RiTerminalWindowLine from "~icons/ri/terminal-window-line?raw";
import RiCheckboxCircleLine from "~icons/ri/checkbox-circle-line?raw";
import RiBarChartHorizontalLine from "~icons/ri/bar-chart-horizontal-line?raw";

const icons = [
  // Element Plus Icon: https://github.com/element-plus/element-plus-icons
  ["ep/home-filled", EpHomeFilled],
  // Remix Icon: https://github.com/Remix-Design/RemixIcon
  ["ri/search-line", RiSearchLine],
  ["ri/information-line", RiInformationLine],
  ["ri/information-line", RiInformationLine],
  ["ri/bill-fill", Bill],
  ["ri/team-line", Team],
  // Element Plus Icon: https://github.com/element-plus/element-plus-icons
  ["ep/menu", EpMenu],
  ["ep/edit", EpEdit],
  ["ep/guide", EpGuide],
  ["ep/set-up", EpSetUp],
  ["ep/monitor", EpMonitor],
  ["ep/lollipop", EpLollipop],
  ["ep/histogram", EpHistogram],
  ["ep/home-filled", EpHomeFilled],
  // Remix Icon: https://github.com/Remix-Design/RemixIcon
  ["ri/mind-map", RiMindMap],
  ["ri/admin-fill", RiAdminFill],
  ["ri/table-line", RiTableLine],
  ["ri/links-fill", RiLinksFill],
  ["ri/admin-line", RiAdminLine],
  ["ri/list-check", RiListCheck],
  ["ri/search-line", RiSearchLine],
  ["ri/window-line", RiWindowLine],
  ["ri/ubuntu-fill", RiUbuntuFill],
  ["ri/history-fill", RiHistoryFill],
  ["ri/edit-box-line", RiEditBoxLine],
  ["ri/code-box-line", RiCodeBoxLine],
  ["ri/artboard-line", RiArtboardLine],
  ["ri/markdown-line", RiMarkdownLine],
  ["ri/file-info-line", RiFileInfoLine],
  ["ri/bank-card-line", RiBankCardLine],
  ["ri/file-ppt-2-line", RiFilePpt2Line],
  ["ri/git-branch-line", RiGitBranchLine],
  ["ri/settings-3-line", RiSettings3Line],
  ["ri/user-voice-line", RiUserVoiceLine],
  ["ri/bookmark-2-line", RiBookmark2Line],
  ["ri/file-search-line", RiFileSearchLine],
  ["ri/chat-search-line", RiChatSearchLine],
  ["ri/information-line", RiInformationLine],
  ["ri/terminal-window-line", RiTerminalWindowLine],
  ["ri/checkbox-circle-line", RiCheckboxCircleLine],
  ["ri/bar-chart-horizontal-line", RiBarChartHorizontalLine],
  ["ri/swap-line", SwapLine],
  ["ep/bell", Bell],
  ["ep/bell-filled", BellFill],
  ["ri/openai-line", Ai],
  ["ri/apple-line", Apple],
  ["ri/apps-2-line", App],
  ["ri/film-line", Film],
  ["ri/video-line", Video],
  ["ri/voice-recognition-fill", Voice],
  ["ri/user-settings-line", UserSetting],
  ["ri/device-line", Device],
  ["ri/key-2-line", Key],
  ["ri/function-line", Function],
  ["ri/building-line", Company],
  ["ri/group-line", Group],
  ["ri/command-line", Cmd],
  ["ep/chat-dot-round", Chat],
  ["ep/collection-tag", Tag],
  ["ep/tickets", Tickets],
  ["ri/robot-line", Bot],
  ["ri/notion-line", No],
  ["ri/upload-line", Counter],
  ["ri/chat-history-line", ChatHistory],
  ["ri/sound-module-line", Module],
  ["ri/shield-user-line", ShieldUser],
  ["ri/archive-drawer-line", Drawer],
  ["ep/histogram", Histogram],
  ["ep/edit", Edit],
  ["ep/guide", Guide],
  ["ep/set-up", SetUp],
  ["ep/monitor", Monitor],
  ["ep/lollipop", Lollipop],
  ["ri/ubuntu-fill", UbuntuFill],
  ["ep/menu", Menu],
  ["ri/ai-generate", Generate],
  ["ri/bank-card-line", Card],
  ["ri/file-ppt-2-line", Ppt],
  ["ep/aim", Aim],
  ["ri/terminal-window-line",TerminalWindowLine],
  ["ri/list-check", ListCheck],
  ["ri/checkbox-circle-line",CheckboxCircleLine],
  ["ri/admin-line", FlUser],
  ["ri/admin-fill", Role],
  ["ri/settings-3-line",Setting],
  ["ri/git-branch-line", Dept],
  ["ri/user-star-line", Customer],
  ["ri/planet-line",Planet],
  ["ri/task-line",Task],
  ["ri/calendar-check-line",CalendarCheck],
  ["ri/open-arm-line",OpenArm],
  ["ri/table-line",Table],
  ["ep/operation",Operation]
];

// 本地菜单图标，后端在路由的 icon 中返回对应的图标字符串并且前端在此处使用 addIcon 添加即可渲染菜单图标
icons.forEach(([name, icon]) => {
  addIcon(name as string, getSvgInfo(icon as string));
});
