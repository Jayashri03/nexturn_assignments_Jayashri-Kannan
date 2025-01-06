class Customer:
    def __init__(self, full_name, contact_email, phone_number):
        self.full_name = full_name
        self.contact_email = contact_email
        self.phone_number = phone_number

    def display_details(self):
        return f"Name: {self.full_name}, Email: {self.contact_email}, Phone: {self.phone_number}"

class CustomerManager:
    def __init__(self, clients=None):
        self.clients = clients if clients else []

    def register_client(self, full_name, contact_email, phone_number):
        customer = Customer(full_name, contact_email, phone_number)
        self.clients.append(customer)
        print("Customer registered successfully!")

    def list_clients(self):
        if not self.clients:
            print("No clients available.")
        else:
            for client in self.clients:
                print(client.display_details())

    def find_client_by_phone(self, phone_number):
        for client in self.clients:
            if phone_number == client.phone_number:
                return client
        return None
