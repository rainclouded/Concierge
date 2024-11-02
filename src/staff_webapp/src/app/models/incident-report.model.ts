export interface IIncidentReport {
	id?: number;
	title: string;
	description: string;
	filing_person_id: number;
	reviewer_id: number;
	severity: 'LOW' | 'MEDIUM' | 'HIGH';
	status: 'OPEN' | 'IN_PROGRESS' | 'RESOLVED' | 'CLOSED';
	created_at: string;
	updated_at: string;
}
