import {ClientReady} from "./events/ClientReady"
import {GuildMemberAdd} from "./events/GuildMemberAdd"
import {GuildMemberRemove} from "./events/GuildMemberRemove"
import {GuildScheduledEventCreate} from "./events/GuildScheduledEventCreate"
import {GuildScheduledEventDelete} from "./events/GuildScheduledEventDelete"
import {GuildScheduledEventUpdate} from "./events/GuildScheduledEventUpdate"
import {InteractionCreate} from "./events/InteractionCreate"
import {MessageCreate} from "./events/MessageCreate"
import {MessageReactionAdd} from "./events/MessageReactionAdd"
import {MessageReactionRemove} from "./events/MessageReactionRemove"
import {ThreadCreate} from "./events/ThreadCreate"
import {ThreadDelete} from "./events/ThreadDelete"
import {PigGrant} from "./commands/PigGrant"
import {PigShow} from "./commands/PigShow"
import {PigPresent} from "./commands/PigPresent";
import {SetNotifyChannel} from "./commands/SetNotifyChannel";
import {StatusGenerate} from "./commands/StatusGenerate";
import {StatusShow} from "./commands/StatusShow";
import {StatusUpdate} from "./commands/StatusUpdate";
import {MeigenAdd} from "./commands/MeigenAdd";
import {MeigenList} from "./commands/MeigenList";
import {MeigenRemove} from "./commands/MeigenRemove";

export const events: Array<any> = [
  ClientReady,
  GuildMemberAdd,
  GuildMemberRemove,
  GuildScheduledEventCreate,
  GuildScheduledEventDelete,
  GuildScheduledEventUpdate,
  InteractionCreate,
  MessageCreate,
  MessageReactionAdd,
  MessageReactionRemove,
  ThreadCreate,
  ThreadDelete
]

export const commands: Array<any> = [
  PigGrant,
  PigShow,
  PigPresent,
  SetNotifyChannel,
  StatusGenerate,
  StatusShow,
  StatusUpdate,
  MeigenAdd,
  MeigenList,
  MeigenRemove
]