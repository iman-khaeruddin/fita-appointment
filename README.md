
## API Reference

#### Create Appointment

```http
  POST /create-appointment
```

| Parameter | Type     | Description               |
| :-------- | :------- |:------------------------- |
| `userId`  | `int`    |                           |
| `coachId` | `int`    |                           |
| `date`    | `string` | `YYYY-MM-DDTHH:mm:ssZXXX` |

#### Coach Decline Appointment

```http
  POST /coach-decline-appointment
```

| Parameter | Type     | Description               |
| :-------- | :------- |:------------------------- |
| `coachId`       | `int`    |                           |
| `appointmentId` | `int`    |                           |

#### Coach Reschedule Appointment

```http
  POST /coach-reschedule-appointment
```
| Parameter | Type     | Description               |
| :-------- | :------- |:------------------------- |
| `coachId`       | `int`    |                           |
| `appointmentId` | `int`    |                           |
| `newDate` | `string`  |`YYYY-MM-DDTHH:mm:ssZXXX`                         |

#### User Decline Appointment

```http
  POST /user-decline-appointment
```
| Parameter | Type     | Description               |
| :-------- | :------- |:------------------------- |
| `userId`       | `int`    |                           |
| `appointmentId` | `int`    |                           |


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`CONNECTION_STRING`


## Flow

- No authentication
- User will book one coach with local timezone
- In BE side will convert datetime and validate based on coach timezone
- If coach available will store the appointment

