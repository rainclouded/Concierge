import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { WindowComponent } from '../../components/window/window.component';
import { AmenityFormComponent } from "../../components/amenity-form/amenity-form.component";
import { IAmenity } from '../../models/amenity.model';
import { AmenityService } from '../../services/amenity.service';
import { ConfirmationDialogComponent } from "../../components/confirmation-dialog/confirmation-dialog.component";
import { ToastrService } from 'ngx-toastr';
import { SessionService } from '../../services/session.service';

@Component({
  selector: 'app-amenities-tab',
  standalone: true,
  imports: [WindowComponent, AmenityFormComponent, ConfirmationDialogComponent, CommonModule],
  templateUrl: './amenities-tab.component.html',
})
export class AmenitiesTabComponent implements OnInit {
  isAmenityWindowOpen = false;
  isConfirmWindowOpen = false;
  amenityToDelete: number | null = null;
  amenities: IAmenity[] = [];
  amenity !: IAmenity;

  sessionPermissionList: string[] | null = null;
  canEdit: boolean = false;
  canDelete: boolean = false;

  constructor(
    private amenityService: AmenityService,
    private sessionService: SessionService,
    private toastr: ToastrService
  ) { }

  ngOnInit(): void {
    this.getAllAmenities();
    this.sessionService.getSessionMe().subscribe(() => {
      this.sessionPermissionList = this.sessionService.sessionPermissionList;
      this.checkPermissions();
    });
  }

  checkPermissions():void {
    if (this.sessionPermissionList) {
      this.canEdit = this.sessionPermissionList.includes('canEditAmenities');
      this.canDelete = this.sessionPermissionList.includes('canDeleteAmenities');
    }
  }

  getAllAmenities() {
    this.amenityService.getAllAmenities()
      .subscribe({
        next: (response) => {
          if (response.data) {
            this.amenities = response.data;
          }
        }
      });
  }

  // load current amenity onto AmenityForm to edit
  loadAmenity(amenity: IAmenity) {
    this.amenity = amenity;
    this.openAmenityWindow();
  }

  deleteAmenity() {
    if (this.amenityToDelete !== null) { // Ensure amenityToDelete is not null
      this.amenityService.deleteAmenity(this.amenityToDelete)
        .subscribe({
          next: (response) => {
            this.getAllAmenities(); // Refresh the amenities list
            console.log(response.message);
            this.toastr.success('Amenity deleted successfully!', 'Delete Successful');
            this.closeConfirmWindow(); // Close the confirmation window
          },
          error: (error) => {
            console.error('Error deleting amenity:', error);
            this.toastr.error('Error deleting amenity!', 'Delete Failed');
          }
        });
    } else {
      console.error('Amenity ID is null, cannot delete.');
    }
  }  

  openAmenityWindow() {
    this.isAmenityWindowOpen = true;
  }

  closeAmenityWindow() {
    this.isAmenityWindowOpen = false;
    this.getAllAmenities();
  }

  openConfirmWindow(amenityId: number) {
    this.isConfirmWindowOpen = true;
    this.amenityToDelete = amenityId;
  }

  closeConfirmWindow() {
    this.isConfirmWindowOpen = false;
    this.getAllAmenities();
  }

  formatTimeToAMPM(time: string): string {
    if (!time) return '';
    const [hours, minutes] = time.split(':');
    const date = new Date();
    date.setHours(parseInt(hours, 10), parseInt(minutes, 10));

    return date.toLocaleTimeString('en-US', {
      hour: '2-digit',
      minute: '2-digit',
      hour12: true
    });
  }

}