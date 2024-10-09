import { Component, OnInit } from '@angular/core';
import { WindowComponent } from '../../components/window/window.component';
import { AmenityFormComponent } from "../../components/amenity-form/amenity-form.component";
import { IAmenity } from '../../models/amenity.model';
import { AmenityService } from '../../services/amenity.service';

@Component({
  selector: 'app-amenities-tab',
  standalone: true,
  imports: [WindowComponent, AmenityFormComponent],
  templateUrl: './amenities-tab.component.html',
})
export class AmenitiesTabComponent implements OnInit {
  isOpenWIndow = false;
  amenities: IAmenity[] = [];
  amenity !: IAmenity;

  constructor(
    private amenityService: AmenityService
  ) { }

  ngOnInit(): void {
    this.getAllAmenities();
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
    this.openWindow();
  }

  deleteAmenity(id: number) {
    this.amenityService.deleteAmenity(id)
      .subscribe({
        next: (response) => {
          this.getAllAmenities();
          console.log(response.message);
        }
      });
  }

  openWindow() {
    this.isOpenWIndow = true;
  }

  closeWindow() {
    this.isOpenWIndow = false;
    this.getAllAmenities();
  }
}