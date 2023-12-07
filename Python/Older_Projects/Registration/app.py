import kivy
from kivy.app import App
from kivy.uix.boxlayout import BoxLayout
from kivy.uix.label import Label
from kivy.uix.textinput import TextInput
from kivy.uix.button import Button
from kivy.uix.popup import Popup
from datetime import datetime
import re
import bcrypt


class RegistrationApp(App):
    def build(self):

        # Box label
        self.title = "Registration Form"
        layout = BoxLayout(orientation="vertical", padding=30, spacing=10)

        # Box title
        head_label = Label(text="Python User Registration App",
                           font_size=26, bold=True, height=40)

        # Adding label
        name_label = Label(text="Name:", font_size=18)
        self.name_input = TextInput(
            multiline=False, font_size=18, input_type="text", halign="center")

        email_label = Label(text="Email:", font_size=18)
        self.email_input = TextInput(
            multiline=False, font_size=18, input_type="text", halign="center")

        password_label = Label(text="Password:", font_size=18)
        self.password_input = TextInput(
            multiline=False, font_size=18, password=True, halign="center")

        confirm_label = Label(text="Confirm password:", font_size=18)
        self.confirm_input = TextInput(
            multiline=False, font_size=18, password=True, halign="center")

        # Add button
        submit_button = Button(
            text="Submit", on_press=self.register, bold=True, font_size=18, halign="center")

        # Adding widgets
        layout.add_widget(head_label)

        layout.add_widget(name_label)
        layout.add_widget(self.name_input)

        layout.add_widget(email_label)
        layout.add_widget(self.email_input)

        layout.add_widget(password_label)
        layout.add_widget(self.password_input)

        layout.add_widget(confirm_label)
        layout.add_widget(self.confirm_input)

        layout.add_widget(submit_button)

        return layout

    def register(self, instance):
        # Collect information
        name = self.name_input.text
        email = self.email_input.text
        password = self.password_input.text
        confirm = self.confirm_input.text

        # Validation
        message = ""

        if name.strip() == "" or email.strip() == "" or password.strip() == "" or confirm.strip() == "":
            message = "Please fill in all fields!"

        elif not is_valid_email(email):
            message = "Invalid email!"

        elif password.strip() != confirm.strip():
            message = "Password do not match!"

        else:
            current_datetime = datetime.now()
            formatted_datetime = current_datetime.strftime(
                "%Y-%m-%d_%H-%M-%S")
            modified_name = name.replace(" ", "-")
            hashed_password = hash_password(password)

            filename = modified_name + "_" + formatted_datetime + ".txt"
            with open(filename, "w") as file:
                file.write('Name: {}\n'.format(name))
                file.write('Email: {}\n'.format(email))
                file.write('Hashed password: {}\n'.format(hashed_password))

            message = "Registration sucessful! Thank you!\n\nDetails you entered:\nName: {}\nEmail: {}".format(
                name, email)

        # Pop-up message
        popup = Popup(title="Registration Status", content=Label(
            text=message), size_hint=(None, None), size=(400, 200))
        popup.open()


def is_valid_email(email):
    pattern = r'^\S+@\S+\.\S+$'
    return bool(re.match(pattern, email))


def hash_password(password):
    # Hash a password for the first time
    hashed_password = bcrypt.hashpw(password.encode('utf-8'), bcrypt.gensalt())
    return hashed_password


if __name__ == "__main__":
    RegistrationApp().run()
