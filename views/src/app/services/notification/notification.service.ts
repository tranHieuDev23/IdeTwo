import { EventEmitter, Injectable, Output } from '@angular/core';
import { IdService } from '../id/id.service';

export enum NotificationType {
  Success,
  Error,
  Info,
  Warning,
}

export class Notification {
  constructor(
    public id: string,
    public type: NotificationType = null,
    public header: string = '',
    public body: string = ''
  ) {}
}

const FIVE_SECOND = 5000;

@Injectable({
  providedIn: 'root',
})
export class NotificationService {
  private notificationList: Notification[] = [];

  @Output() public notificationListChange = new EventEmitter<Notification[]>();

  constructor(private readonly idService: IdService) {}

  public createNotification(
    type: NotificationType,
    header: string = '',
    body: string = ''
  ): void {
    const id = this.idService.makeId();
    this.notificationList = [
      ...this.notificationList,
      new Notification(id, type, header, body),
    ];
    this.notificationListChange.emit(this.notificationList);
    setTimeout(() => {
      this.notificationList = this.notificationList.filter(
        (item) => item.id !== id
      );
      this.notificationListChange.emit(this.notificationList);
    }, FIVE_SECOND);
  }

  public successNotification(header: string = '', body: string = ''): void {
    this.createNotification(NotificationType.Success, header, body);
  }

  public errorNotification(header: string = '', body: string = ''): void {
    this.createNotification(NotificationType.Error, header, body);
  }

  public infoNotification(header: string = '', body: string = ''): void {
    this.createNotification(NotificationType.Info, header, body);
  }

  public warningNotification(header: string = '', body: string = ''): void {
    this.createNotification(NotificationType.Warning, header, body);
  }
}
