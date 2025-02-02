import unittest
import requests

BASE_URL = "http://localhost:8080/api/v1"

class ProjectManagementSystemTest(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        cls.client = requests.Session()

        # Регистрация и авторизация пользователей
        cls.manager_token = cls.register_and_login("managerUser", "password")
        cls.developer_token = cls.register_and_login("developerUser", "password")
        cls.tester_token = cls.register_and_login("testerUser", "password")
        cls.team_leader_token = cls.register_and_login("teamLeaderUser", "password")
        cls.other_user_token = cls.register_and_login("otherUser", "password")

    @classmethod
    def register_and_login(cls, username, password):
        # Регистрация
        register_body = {
            "username": username,
            "password": password,
            "email": f"{username}@example.com"
        }
        register_response = cls.client.post(f"{BASE_URL}/users/register", json=register_body)
        assert register_response.status_code == 201, "User should be registered successfully"

        # Авторизация
        login_body = {
            "username": username,
            "password": password
        }
        login_response = cls.client.post(f"{BASE_URL}/users/login", json=login_body)
        assert login_response.status_code == 200, "User should be logged in successfully"

        return login_response.json().get("token")

    def test_1_create_project(self):
        # Создание проекта менеджером
        project_data = {
            "projectName": "New Project",
            "description": "Project Description"
        }
        response = self.client.post(
            f"{BASE_URL}/projects",
            json=project_data,
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertEqual(response.status_code, 201, "Project should be created successfully")
        self.project_id = response.json().get("projectId")
        self.assertIsNotNone(self.project_id, "Project ID should not be null")

    def test_2_assign_team_leader(self):
        # Назначение тимлидера
        assign_data = {
            "userId": "teamLeaderUser"
        }
        response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/teamleader",
            json=assign_data,
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertEqual(response.status_code, 200, "Team leader should be assigned successfully")

    def test_3_add_developer_and_tester(self):
        # Добавление разработчика
        dev_data = {
            "userId": "developerUser"
        }
        dev_response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/developers",
            json=dev_data,
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertEqual(dev_response.status_code, 200, "Developer should be added successfully")

        # Добавление тестировщика
        tester_data = {
            "userId": "testerUser"
        }
        tester_response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/testers",
            json=tester_data,
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertEqual(tester_response.status_code, 200, "Tester should be added successfully")

    def test_4_create_milestone(self):
        # Создание майлстоуна
        milestone_data = {
            "name": "Milestone 1",
            "startDate": "2024-01-01",
            "endDate": "2024-02-01"
        }
        response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/milestones",
            json=milestone_data,
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertEqual(response.status_code, 201, "Milestone should be created successfully")
        self.milestone_id = response.json().get("milestoneId")
        self.assertIsNotNone(self.milestone_id, "Milestone ID should not be null")

    def test_5_create_and_assign_ticket(self):
        # Создание тикета
        ticket_data = {
            "title": "Implement feature X",
            "description": "Details about feature X"
        }
        response = self.client.post(
            f"{BASE_URL}/milestones/{self.milestone_id}/tickets",
            json=ticket_data,
            headers={"Authorization": f"Bearer {self.team_leader_token}"}
        )
        self.assertEqual(response.status_code, 201, "Ticket should be created successfully")
        self.ticket_id = response.json().get("ticketId")
        self.assertIsNotNone(self.ticket_id, "Ticket ID should not be null")

        # Назначение тикета разработчику
        assign_data = {
            "userId": "developerUser"
        }
        assign_response = self.client.post(
            f"{BASE_URL}/tickets/{self.ticket_id}/assign",
            json=assign_data,
            headers={"Authorization": f"Bearer {self.team_leader_token}"}
        )
        self.assertEqual(assign_response.status_code, 200, "Ticket should be assigned successfully")

    def test_6_update_ticket_status(self):
        # Обновление статуса тикета
        status_data = {"status": "in_progress"}
        response = self.client.put(
            f"{BASE_URL}/tickets/{self.ticket_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.developer_token}"}
        )
        self.assertEqual(response.status_code, 200, "Ticket status should be updated to in_progress")

        status_data = {"status": "completed"}
        response = self.client.put(
            f"{BASE_URL}/tickets/{self.ticket_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.developer_token}"}
        )
        self.assertEqual(response.status_code, 200, "Ticket status should be updated to completed")

    def test_7_check_ticket_completion(self):
        # Проверка статуса тикета
        response = self.client.get(
            f"{BASE_URL}/tickets/{self.ticket_id}/status",
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertEqual(response.status_code, 200, "Should retrieve ticket status successfully")
        self.assertEqual(response.json().get("status"), "completed", "Ticket status should be 'completed'")

    def test_8_create_bug_report(self):
        # Тестирование проекта
        response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/test",
            headers={"Authorization": f"Bearer {self.tester_token}"}
        )
        self.assertEqual(response.status_code, 200, "Project should be tested successfully")

        # Создание баг-репорта
        bug_report_data = {
            "description": "Found a bug in feature X"
        }
        response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/bugreports",
            json=bug_report_data,
            headers={"Authorization": f"Bearer {self.tester_token}"}
        )
        self.assertEqual(response.status_code, 201, "Bug report should be created successfully")
        self.bug_report_id = response.json().get("bugReportId")
        self.assertIsNotNone(self.bug_report_id, "Bug report ID should not be null")

    def test_9_fix_bug_report(self):
        # Исправление баг-репорта
        status_data = {"status": "fixed"}
        response = self.client.put(
            f"{BASE_URL}/bugreports/{self.bug_report_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.developer_token}"}
        )
        self.assertEqual(response.status_code, 200, "Bug report status should be updated to 'fixed'")

    def test_10_verify_bug_fix(self):
        # Проверка исправления баг-репорта
        status_data = {"status": "tested"}
        response = self.client.put(
            f"{BASE_URL}/bugreports/{self.bug_report_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.tester_token}"}
        )
        self.assertEqual(response.status_code, 200, "Bug report status should be updated to 'tested'")

        status_data = {"status": "closed"}
        response = self.client.put(
            f"{BASE_URL}/bugreports/{self.bug_report_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.tester_token}"}
        )
        self.assertEqual(response.status_code, 200, "Bug report status should be updated to 'closed'")

    def test_11_close_milestone(self):
        # Закрытие майлстоуна
        status_data = {"status": "closed"}
        response = self.client.put(
            f"{BASE_URL}/milestones/{self.milestone_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertEqual(response.status_code, 200, "Milestone status should be updated to 'closed'")

    def test_12_developer_assign_ticket(self):
        # Попытка разработчика назначить тикет
        assign_data = {
            "userId": "otherUser"
        }
        response = self.client.post(
            f"{BASE_URL}/tickets/{self.ticket_id}/assign",
            json=assign_data,
            headers={"Authorization": f"Bearer {self.developer_token}"}
        )
        self.assertNotEqual(response.status_code, 200, "Developer should not be able to assign tickets")

    def test_13_tester_create_milestone(self):
        # Попытка тестировщика создать майлстоун
        milestone_data = {
            "name": "Invalid Milestone",
            "startDate": "2024-03-01",
            "endDate": "2024-04-01"
        }
        response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/milestones",
            json=milestone_data,
            headers={"Authorization": f"Bearer {self.tester_token}"}
        )
        self.assertEqual(response.status_code, 403, "Tester should not be able to create milestones")

    def test_14_access_nonexistent_project(self):
        # Попытка доступа к несуществующему проекту
        response = self.client.post(
            f"{BASE_URL}/projects/InvalidId/test",
            headers={"Authorization": f"Bearer {self.tester_token}"}
        )
        self.assertEqual(response.status_code, 404, "Should receive 404 Not Found for nonexistent project")

    def test_15_access_project_unauthorized(self):
        # Попытка доступа к проекту без прав
        response = self.client.post(
            f"{BASE_URL}/projects/{self.project_id}/test",
            headers={"Authorization": f"Bearer {self.other_user_token}"}
        )
        self.assertEqual(response.status_code, 403, "User should not have access to the project")

    def test_16_update_bug_report_without_access(self):
        # Попытка обновить баг-репорт без прав
        status_data = {"status": "closed"}
        response = self.client.put(
            f"{BASE_URL}/bugreports/{self.bug_report_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.other_user_token}"}
        )
        self.assertNotEqual(response.status_code, 200, "User should not be able to update bug reports they do not have access to")

    def test_17_create_ticket_in_closed_milestone(self):
        # Попытка создать тикет в закрытом майлстоуне
        ticket_data = {
            "title": "New Task",
            "description": "Details"
        }
        response = self.client.post(
            f"{BASE_URL}/milestones/{self.milestone_id}/tickets",
            json=ticket_data,
            headers={"Authorization": f"Bearer {self.team_leader_token}"}
        )
        self.assertEqual(response.status_code, 403, "Should not be able to create ticket in a closed milestone")

    def test_18_manager_perform_developer_action(self):
        # Попытка менеджера выполнить действие разработчика
        status_data = {"status": "fixed"}
        response = self.client.put(
            f"{BASE_URL}/bugreports/{self.bug_report_id}/status",
            json=status_data,
            headers={"Authorization": f"Bearer {self.manager_token}"}
        )
        self.assertNotEqual(response.status_code, 200, "Manager should not be able to perform developer actions")

if __name__ == "__main__":
    unittest.main()