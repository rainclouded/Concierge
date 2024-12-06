export interface ISessionData {
	sessionData: IDataInfo;
}

export interface IDataInfo {
	SessionPermissionList: string[];
	accountId: number;
	accountName: string;
}