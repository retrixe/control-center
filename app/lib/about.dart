import 'package:control_center/page.dart';
import 'package:dbus/dbus.dart';
import 'package:flutter/material.dart';
import 'package:url_launcher/url_launcher.dart';

class AboutPage extends SettingsPage {
  const AboutPage({Key? key, required DBusClient client})
      : super(key: key, title: 'About', client: client);

  @override
  Widget buildPage(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: SettingCategory(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: [
            Text("About", style: Theme.of(context).textTheme.headline5),
            const Padding(padding: EdgeInsets.symmetric(vertical: 8.0)),
            const Text("A Linux app for settings usually not in your desktop's "
                "Settings app. Made of a Flutter desktop application and a "
                "daemon written in Golang.\n"),
            InkWell(
                child: const Text(
                    "Built with <3 at https://github.com/retrixe/control-center"),
                onTap: () => launchUrl(
                      Uri.parse("https://github.com/retrixe/control-center"),
                    )),
            const Text("\nOther projects to look at:"),
            InkWell(
                child:
                    const Text("- OpenRazer/Polychromatic for Razer devices"),
                onTap: () => launchUrl(
                      Uri.parse("https://polychromatic.app/"),
                    )),
            InkWell(
              child: const Text("- Piper for mice and keyboards"),
              onTap: () => launchUrl(Uri.parse(
                  "https://github.com/libratbag/piper/wiki/Installation")),
            ),
            InkWell(
              child: const Text("- asusctl for ASUS ROG laptops"),
              onTap: () =>
                  launchUrl(Uri.parse("https://gitlab.com/asus-linux/asusctl")),
            ),
          ],
        ),
      ),
    );
  }
}
