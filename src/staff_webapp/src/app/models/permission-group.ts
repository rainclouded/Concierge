import { IPermission } from "./permission";

export interface IPermissionGroup {
  groupId: number;
  groupName: string;
  groupDescription: string;
  groupPermissions: IPermission[];
  groupMembers: number[];
}
